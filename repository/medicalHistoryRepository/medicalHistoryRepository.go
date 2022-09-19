package medicalHistoryRepository

import (
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/requests"
	"github.com/saeedhpro/apisimateb/repository"
)

func GetUserMedicalHistory(id uint64) (*models.MedicalHistoryOrthodontics, error) {
	history := models.MedicalHistoryOrthodontics{UserID: id}
	err := repository.DB.MySQL.First(&history, &history).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func CreateUserMedicalHistory(request *requests.CreateUserMedicalHistoryRequest) error {
	history := models.MedicalHistoryOrthodontics{
		UserID:                  request.UserID,
		AdenoidTonsileReduction: request.AdenoidTonsileReduction,
		MedicalCondition:        request.MedicalCondition,
		ConsumableMedicine:      request.ConsumableMedicine,
		GeneralHealth:           request.GeneralHealth,
		UnderPhysicianCare:      request.UnderPhysicianCare,
		AccidentToHead:          request.AccidentToHead,
		Operations:              request.Operations,
		ChiefComplaint:          request.ChiefComplaint,
		PreviousOrthodontic:     request.PreviousOrthodontic,
		OralHygiene:             request.OralHygiene,
		Frontal:                 request.Frontal,
		Profile:                 request.Profile,
		TeethPresent:            request.TeethPresent,
		UnErupted:               request.UnErupted,
		IeMissing:               request.IeMissing,
		IeExtracted:             request.IeExtracted,
		IeImpacted:              request.IeImpacted,
		IeSupernumerary:         request.IeSupernumerary,
		IeCaries:                request.IeCaries,
		IeRct:                   request.IeRct,
		IeAnomalies:             request.IeAnomalies,
		LeftMolar:               request.LeftMolar,
		RightMolar:              request.RightMolar,
		LeftCanine:              request.LeftCanine,
		RightCanine:             request.RightCanine,
		Overjet:                 request.Overjet,
		Overbite:                request.Overbite,
		Crossbite:               request.Crossbite,
		CrowdingMd:              request.CrowdingMd,
		CrowdingMx:              request.CrowdingMx,
		SpacingMx:               request.SpacingMx,
		SpacingMd:               request.SpacingMd,
		Diagnosis:               request.Diagnosis,
		TreatmentPlan:           request.TreatmentPlan,
		LengthActiveTreatment:   request.LengthActiveTreatment,
		Retention:               request.Retention,
	}
	err := repository.DB.MySQL.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserMedicalHistory(request *requests.CreateUserMedicalHistoryRequest) error {
	history := models.MedicalHistoryOrthodontics{
		UserID:                  request.UserID,
		AdenoidTonsileReduction: request.AdenoidTonsileReduction,
		MedicalCondition:        request.MedicalCondition,
		ConsumableMedicine:      request.ConsumableMedicine,
		GeneralHealth:           request.GeneralHealth,
		UnderPhysicianCare:      request.UnderPhysicianCare,
		AccidentToHead:          request.AccidentToHead,
		Operations:              request.Operations,
		ChiefComplaint:          request.ChiefComplaint,
		PreviousOrthodontic:     request.PreviousOrthodontic,
		OralHygiene:             request.OralHygiene,
		Frontal:                 request.Frontal,
		Profile:                 request.Profile,
		TeethPresent:            request.TeethPresent,
		UnErupted:               request.UnErupted,
		IeMissing:               request.IeMissing,
		IeExtracted:             request.IeExtracted,
		IeImpacted:              request.IeImpacted,
		IeSupernumerary:         request.IeSupernumerary,
		IeCaries:                request.IeCaries,
		IeRct:                   request.IeRct,
		IeAnomalies:             request.IeAnomalies,
		LeftMolar:               request.LeftMolar,
		RightMolar:              request.RightMolar,
		LeftCanine:              request.LeftCanine,
		RightCanine:             request.RightCanine,
		Overjet:                 request.Overjet,
		Overbite:                request.Overbite,
		Crossbite:               request.Crossbite,
		CrowdingMd:              request.CrowdingMd,
		CrowdingMx:              request.CrowdingMx,
		SpacingMx:               request.SpacingMx,
		SpacingMd:               request.SpacingMd,
		Diagnosis:               request.Diagnosis,
		TreatmentPlan:           request.TreatmentPlan,
		LengthActiveTreatment:   request.LengthActiveTreatment,
		Retention:               request.Retention,
	}
	err := repository.DB.MySQL.
		Model(&history).
		Where("user_id = ?", &history.UserID).
		Updates(history).
		Error
	if err != nil {
		return err
	}
	return nil
}
