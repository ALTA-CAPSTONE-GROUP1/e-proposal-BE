package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin"
	"github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/submission"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type submissionModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) submission.Repository {
	return &submissionModel{
		db: db,
	}
}

func (sm *submissionModel) FindRequirement(userID string, typeName string, typeValue int) (submission.Core, error) {
	var applicant admin.Users
	var typeDetail admin.Type
	var tos []admin.Users
	var ccs []admin.Users

	applicant.ID = userID

	if err := sm.db.Where("name = ? ", typeName).Find(&typeDetail).Error; err != nil {
		log.Errorf("error on finding typeDetails from %s", typeName)
		return submission.Core{}, err
	}

	if err := sm.db.Preload("Office").First(&applicant).Error; err != nil {
		log.Errorf("error on finding applicant office %s", applicant.Name)
		return submission.Core{}, err
	}

	if err := sm.db.Raw(`
    SELECT users.id, users.name, users.email, users.phone_number, offices.id as office_id, position_has_types.as AS position_role
    FROM users
    INNER JOIN position_has_types ON position_has_types.position_id = users.position_id
    INNER JOIN positions ON positions.id = position_has_types.position_id
    INNER JOIN offices ON offices.id = users.office_id
    INNER JOIN types ON types.id = position_has_types.type_id
    WHERE positions.deleted_at IS NULL AND types.deleted_at IS NULL AND position_has_types.deleted_at IS NULL
    AND types.name = ? AND position_has_types.as = 'cc' AND position_has_types.value = ? AND (users.office_id = ? OR offices.name = 'Head Office')
`, typeName, 30000000, applicant.Office.ID).Scan(&ccs).Error; err != nil {
		return submission.Core{}, err
	}

	if err := sm.db.Raw(`
    SELECT users.id, users.name, users.email, users.phone_number, offices.id as office_id, position_has_types.as AS position_role
    FROM users
    INNER JOIN position_has_types ON position_has_types.position_id = users.position_id
    INNER JOIN positions ON positions.id = position_has_types.position_id
    INNER JOIN offices ON offices.id = users.office_id
    INNER JOIN types ON types.id = position_has_types.type_id
    WHERE positions.deleted_at IS NULL AND types.deleted_at IS NULL AND position_has_types.deleted_at IS NULL
    AND types.name = ? AND position_has_types.as = 'to' AND position_has_types.value = ? AND (users.office_id = ? OR offices.name = 'Head Office')
	ORDER BY position_has_types.to_level ASC
`, typeName, 30000000, applicant.Office.ID).Scan(&tos).Error; err != nil {
		return submission.Core{}, err
	}

	var result submission.Core

	result.Requirement = typeDetail.Requirement

	for _, to := range tos {
		tmp := submission.ToApprover{
			ApproverPosition: to.Position.Name,
			ApproverId:       to.ID,
			ApproverName:     to.Name,
		}
		result.To = append(result.To, tmp)
	}

	for _, cc := range ccs {
		tmp := submission.CcApprover{
			CcPosition: cc.Position.Name,
			CcId:       cc.ID,
			CcName:     cc.Name,
		}
		result.CC = append(result.CC, tmp)
	}

	return result, nil
}

func (sm *submissionModel) InsertSubmission(newSub submission.AddSubmissionCore) error {
	var submissionDB Submission

	var SubmissionTypeQuery admin.Type

	if err := sm.db.Where("name = ?", newSub.SubmissionType).First(&SubmissionTypeQuery).Error; err != nil {
		log.Error("cannot find submissiontype by name")
		return err
	}

	submissionDB.UserID = newSub.OwnerID
	submissionDB.Title = newSub.Title
	submissionDB.TypeID = SubmissionTypeQuery.ID
	submissionDB.Is_Opened = false
	submissionDB.UserID = newSub.OwnerID
	submissionDB.Status = "Sent"

	for _, v := range newSub.ToApprover {
		tmp := To{
			Name:   v.ApproverName,
			UserID: v.ApproverId,
		}
		submissionDB.Tos = append(submissionDB.Tos, tmp)
	}

	file := File{
		Name: newSub.Attachment,
		Link: newSub.AttachmentLink,
	}
	submissionDB.Files = append(submissionDB.Files, file)

	if err := sm.db.Create(&submissionDB).Error; err != nil {
		log.Error("error occurs while insert submission datas")
		return err
	}

	return nil
}

func (sm *submissionModel) SelectAllSubmissions(userID string, pr submission.GetAllQueryParams) ([]submission.AllSubmiisionCore, []admin.Type, error) {
	var (
		dbsubmissions       []Submission
		resultAllSubmission []submission.AllSubmiisionCore
		subTypes            []admin.Type
		user                admin.Users
	)

	if pr.Limit < 1 {
		pr.Limit = 10
	}
	if pr.Offset < 0 {
		pr.Offset = 0
	}

	if err := sm.db.Where("id = ?", userID).Preload("Position.Types").Find(&user).Error; err != nil {
		log.Errorf("error on finding subTypes have by user", err)
		return []submission.AllSubmiisionCore{}, []admin.Type{}, err
	}

	subTypes = append(subTypes, user.Position.Types...)

	if err := sm.db.Where("user_id = ?", userID).Find(&dbsubmissions).Error; err != nil {
		log.Errorf("error on finding submissions for user %s: %v", userID, err)
		return []submission.AllSubmiisionCore{}, []admin.Type{}, err
	}

	for _, sub := range dbsubmissions {
		var toApprover []submission.ToApprover
		for _, to := range sub.Tos {
			var toDetails admin.Users
			if err := sm.db.Where("id = ?", to.UserID).Preload("Positions").Find(&toDetails).Error; err != nil {
				log.Error("failed on finding positions of tos")
			}
			toApprover = append(toApprover, submission.ToApprover{
				ApproverId:       to.UserID,
				ApproverName:     to.Name,
				ApproverPosition: toDetails.Position.Name,
			})
		}
		var ccApprover []submission.CcApprover
		for _, cc := range sub.Ccs {
			var ccDetails admin.Users
			if err := sm.db.Where("id = ?", cc.UserID).Preload("Positions").Find(&ccDetails).Error; err != nil {
				log.Error("failed on finding positions of tos")
			}
			ccApprover = append(ccApprover, submission.CcApprover{
				CcPosition: ccDetails.Position.Name,
				CcName:     cc.Name,
				CcId:       cc.UserID,
			})
		}

		var attachment File
		if err := sm.db.Where("submission_id = ?", sub.ID).First(&attachment).Error; err != nil {
			log.Errorf("error getting files for submission %d: %v", sub.ID, err)
			return []submission.AllSubmiisionCore{}, []admin.Type{}, err
		}

		var subTypeByID admin.Type
		if err := sm.db.Where("id = ?", sub.TypeID).First(&subTypeByID).Error; err != nil {
			log.Errorf("error getting files for subType %d: %v", sub.TypeID, err)
			return []submission.AllSubmiisionCore{}, []admin.Type{}, err
		}

		resultAllSubmission = append(resultAllSubmission, submission.AllSubmiisionCore{
			ID:             sub.ID,
			Tos:            toApprover,
			CCs:            ccApprover,
			Title:          sub.Title,
			Status:         sub.Status,
			ReceiveDate:    sub.CreatedAt.Format(time.RFC3339),
			Opened:         sub.Is_Opened,
			Attachment:     attachment.Link,
			SubmissionType: subTypeByID.Name,
		})
	}
	if pr.To != "" {
		var filteredByTo []submission.AllSubmiisionCore
		for _, data := range resultAllSubmission {
			for _, v := range data.Tos {
				if strings.Contains(strings.ToLower(v.ApproverName), strings.ToLower(pr.To)) ||
					strings.Contains(strings.ToLower(v.ApproverId), strings.ToLower(pr.To)) ||
					strings.Contains(strings.ToLower(v.ApproverPosition), strings.ToLower(pr.To)) {
					filteredByTo = append(filteredByTo, data)
				}
			}
			resultAllSubmission = filteredByTo
		}
	}

	if pr.Title != "" {
		var filteredByTitle []submission.AllSubmiisionCore
		for _, data := range resultAllSubmission {
			if strings.Contains(strings.ToLower(data.Title), strings.ToLower(pr.Title)) {
				filteredByTitle = append(filteredByTitle, data)
			}
		}
		resultAllSubmission = filteredByTitle
	}

	end := pr.Limit + pr.Limit
	if end > len(resultAllSubmission) {
		end = len(resultAllSubmission)
	}

	resultAllSubmission = resultAllSubmission[pr.Offset:end]

	return resultAllSubmission, subTypes, nil
}

func (sm *submissionModel) SelectSubmissionByID(submissionID int) (submission.GetSubmmisionByIDCore, error) {
	var result submission.GetSubmmisionByIDCore

	var submissions Submission
	if err := sm.db.Preload("To").
		Preload("Signs").
		Preload("Ccs").
		Preload("Files").
		First(&submissions, submissionID).Error; err != nil {
		log.Errorf("error on finding all data from submission ID", err)
		return submission.GetSubmmisionByIDCore{},
			fmt.Errorf("error on finding submissionDatas %w", err)
	}

	result.Attachment = submissions.Files[0].Link

	var typeDetails admin.Type
	if err := sm.db.Where("id = ?", submissions.TypeID).
		First(&typeDetails).Error; err != nil {
		log.Errorf("error on finding submissionType detail %w", err)
		return submission.GetSubmmisionByIDCore{},
			fmt.Errorf("error on finding submissionType %w", err)
	}

	result.SubmissionType = typeDetails.Name

	for _, submissionTo := range submissions.Tos {
		tmp := submission.ToApprover{
			ApproverPosition: submissionTo.Name,
			ApproverId:       submissionTo.UserID,
		}
		tmpaction := submission.ApproverActions{
			Action:       submissionTo.Action_Type,
			ApproverName: submissionTo.Name,
			Message:      submissionTo.Message,
		}
		result.To = append(result.To, tmp)
		result.ApproverActions = append(result.ApproverActions, tmpaction)
	}

	for i, v := range result.To {
		var to admin.Users
		if err := sm.db.Preload("Positions").
			Where("id = ?", v.ApproverId).
			First(&to).Error; err != nil {
			log.Errorf("eror on finding to position %w", err)
			return submission.GetSubmmisionByIDCore{},
				fmt.Errorf("error on to positions %w", err)
		}
		result.To[i].ApproverPosition = to.Position.Name
		result.ApproverActions[i].ApproverPosition = to.Position.Name
	}

	for _, submissionCc := range submissions.Ccs {
		tmp := submission.CcApprover{
			CcName: submissionCc.Name,
			CcId:   submissionCc.UserID,
		}

		result.CC = append(result.CC, tmp)
	}

	for i, v := range result.CC {
		var cc admin.Users
		if err := sm.db.Preload("Positions").
			Where("id = ?", v.CcId).
			First(&cc).Error; err != nil {
			log.Errorf("eror on finding to position %w", err)
			return submission.GetSubmmisionByIDCore{},
				fmt.Errorf("error on to positions %w", err)
		}

		result.CC[i].CcPosition = cc.Position.Name
	}

	result.Message = submissions.Message
	result.Title = submissions.Title

	return result, nil
}
