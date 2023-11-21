package db

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/buniekbua/task-managment-system/internal/models"
	"github.com/buniekbua/task-managment-system/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var myTime time.Time = time.Date(2023, time.November, 11, 0, 0, 0, 0, time.UTC)

func TestCreateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	s := NewTaskRepository(dbx)

	tests := []struct {
		name    string
		s       *TaskRepository
		task    *models.Task
		mock    func()
		wantErr error
	}{
		{
			name: "OK",
			s:    s,
			task: &models.Task{
				UserID:      1,
				TaskName:    "Test Task",
				Description: "Test Description",
				DueDate:     "2022-11-11",
				Priority:    "low",
				CreatedAt:   myTime,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"task_id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO tasks").WithArgs(1, "Test Task", "Test Description", util.StrToDate("2022-11-11"), "low", sqlmock.AnyArg()).WillReturnRows(rows)
			},
			wantErr: nil,
		},
		{
			name: "Empty fields",
			s:    s,
			task: &models.Task{
				UserID:      1,
				TaskName:    "",
				Description: "",
				DueDate:     "2022-11-11",
				Priority:    "low",
				CreatedAt:   myTime,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"task_id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO tasks").WithArgs(1, "", "", util.StrToDate("2022-11-11"), "low", sqlmock.AnyArg()).WillReturnRows(rows)
			},
			wantErr: errors.New("empty fields"),
		},
		{
			name: "Wrong priority",
			s:    s,
			task: &models.Task{
				UserID:      1,
				TaskName:    "Test Task",
				Description: "Test Description",
				DueDate:     "2022-11-11",
				Priority:    "wrong",
				CreatedAt:   myTime,
			},
			mock: func() {
				mock.ExpectQuery("INSERT INTO tasks").WithArgs(1, "Test Task", "Test Description", "2022-11-11", "wrong", sqlmock.AnyArg()).WillReturnError(errors.New("invalid priority"))
			},
			wantErr: errors.New("invalid priority"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mock()

			err := tt.s.CreateTask(tt.task)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestGetTaskByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	r := &TaskRepository{db: dbx}

	//myTime := time.Date(2023, time.November, 11, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		r       *TaskRepository
		id      int
		mock    func()
		want    *models.Task
		wantErr bool
	}{
		{
			name: "OK",
			r:    r,
			id:   1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"task_id", "user_id", "task_name", "description", "due_date", "priority", "created_at"}).
					AddRow(1, 1, "Test Task", "Test Description", "2022-11-11", "low", myTime)
				mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)
			},
			want: &models.Task{
				TaskID:      1,
				UserID:      1,
				TaskName:    "Test Task",
				Description: "Test Description",
				DueDate:     "2022-11-11",
				Priority:    "low",
				CreatedAt:   myTime,
			},
			wantErr: false,
		},
		{
			name: "Not found",
			r:    r,
			id:   1,
			mock: func() {
				mock.ExpectQuery("SELECT").WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetTaskByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
