package services

import (
	"fmt"
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/models"
	"meetspace_backend/auth/repositories"
	"meetspace_backend/auth/types"
	"meetspace_backend/utils"
)

type VerificationService struct {
    VerificationRepository *repositories.VerificationRepository
}

func NewVerificationService(repo *repositories.VerificationRepository) *VerificationService {
    return &VerificationService{
        VerificationRepository: repo,
    }
}


func (v *VerificationService) Create(reqData types.SendEmailRequest) *utils.Response {
    // validate request struct data
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
    }

	otp := utils.GenerateOTP()
	
	emailBody := fmt.Sprintf("Your OTP is:- <h2> %s </h2>", otp)
	go utils.SendEmail(reqData.Email, "Email OTP", emailBody)
	
	data := models.Verification{
		Email: reqData.Email,
		Otp: otp,
	}
	v.VerificationRepository.CreateRecord(data)

    return utils.SuccessResponse("success", nil)
}


func (v *VerificationService) VerifyEmailOtp(reqData types.VerifyEmailRequest) *utils.Response {
	// validate request struct data
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
    }

	_, err := v.VerificationRepository.GetRecordByEmailAndOtp(reqData.Email, reqData.OTP)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

	return utils.SuccessResponse("success", nil)
}
