package message

import (
	"bytes"
	"einfach-msg/contract"
	"einfach-msg/model"
	"einfach-msg/test/mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMessage_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	collectionMock := mock.NewMockCollection(ctrl)
	request := model.Model{
		Subject:    "1",
		Message:    "1",
		SenderID:   "1",
		ReceiverID: "1",
	}

	body, _ := json.Marshal(request)
	type fields struct {
		collection contract.Collection
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type mockChecker struct {
		checkInsert bool
	}
	type insertScenario struct {
		err error
	}
	type mockScenario struct {
		mockChecker    mockChecker
		insertScenario insertScenario
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockScenario mockScenario
	}{
		{
			name: "1. negative | cannot read request",
			fields: fields{
				collection: collectionMock,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/v1/message", bytes.NewBuffer(nil)),
			},
		},
		{
			name: "2. negative | error while insert document",
			fields: fields{
				collection: collectionMock,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/v1/message", bytes.NewBuffer(body)),
			},
			mockScenario: mockScenario{
				mockChecker: mockChecker{
					checkInsert: true,
				},
				insertScenario: insertScenario{
					err: errors.New("failed insert"),
				},
			},
		},
		{
			name: "3. positive | successfully sending message",
			fields: fields{
				collection: collectionMock,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/v1/message", bytes.NewBuffer(body)),
			},
			mockScenario: mockScenario{
				mockChecker: mockChecker{
					checkInsert: true,
				},
				insertScenario: insertScenario{
					err: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				collection: tt.fields.collection,
			}
			if tt.mockScenario.mockChecker.checkInsert {
				collectionMock.EXPECT().Insert(gomock.Any()).Return(tt.mockScenario.insertScenario.err)
			}
			m.Create(tt.args.w, tt.args.r)
		})
	}
}