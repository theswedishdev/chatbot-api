package types

type AdviceResponse struct {
	Slip AdviceSlip `json:"slip,omitempty"`
	Message AdviceErrorMessage `json:"message,omitempty"`
}

type AdviceSlip struct {
	ID     int    `json:"id"`
	Advice string `json:"advice"`
}

type AdviceErrorMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (ar *AdviceResponse) GetID() int {
	return ar.Slip.ID
}

func (ar *AdviceResponse) GetAdvice() string {
	return ar.Slip.Advice
}

func (ar *AdviceResponse) GetErrorMessage() string {
	return ar.Message.Text
}

func (ar *AdviceResponse) HasError() bool {
	return ar.Message != AdviceErrorMessage{}
}
