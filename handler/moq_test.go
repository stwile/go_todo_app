// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/stwile/go_todo_app/entity"
	"sync"
)

// Ensure, that ListTasksServiceMock does implement ListTasksService.
// If this is not the case, regenerate this file with moq.
var _ ListTasksService = &ListTasksServiceMock{}

// ListTasksServiceMock is a mock implementation of ListTasksService.
//
//	func TestSomethingThatUsesListTasksService(t *testing.T) {
//
//		// make and configure a mocked ListTasksService
//		mockedListTasksService := &ListTasksServiceMock{
//			ListTasksFunc: func(ctx context.Context) (entity.Tasks, error) {
//				panic("mock out the ListTasks method")
//			},
//		}
//
//		// use mockedListTasksService in code that requires ListTasksService
//		// and then make assertions.
//
//	}
type ListTasksServiceMock struct {
	// ListTasksFunc mocks the ListTasks method.
	ListTasksFunc func(ctx context.Context) (entity.Tasks, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListTasks holds details about calls to the ListTasks method.
		ListTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockListTasks sync.RWMutex
}

// ListTasks calls ListTasksFunc.
func (mock *ListTasksServiceMock) ListTasks(ctx context.Context) (entity.Tasks, error) {
	if mock.ListTasksFunc == nil {
		panic("ListTasksServiceMock.ListTasksFunc: method is nil but ListTasksService.ListTasks was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListTasks.Lock()
	mock.calls.ListTasks = append(mock.calls.ListTasks, callInfo)
	mock.lockListTasks.Unlock()
	return mock.ListTasksFunc(ctx)
}

// ListTasksCalls gets all the calls that were made to ListTasks.
// Check the length with:
//
//	len(mockedListTasksService.ListTasksCalls())
func (mock *ListTasksServiceMock) ListTasksCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListTasks.RLock()
	calls = mock.calls.ListTasks
	mock.lockListTasks.RUnlock()
	return calls
}

// Ensure, that AddTaskServiceMock does implement AddTaskService.
// If this is not the case, regenerate this file with moq.
var _ AddTaskService = &AddTaskServiceMock{}

// AddTaskServiceMock is a mock implementation of AddTaskService.
//
//	func TestSomethingThatUsesAddTaskService(t *testing.T) {
//
//		// make and configure a mocked AddTaskService
//		mockedAddTaskService := &AddTaskServiceMock{
//			AddTaskFunc: func(ctx context.Context, title string) (*entity.Task, error) {
//				panic("mock out the AddTask method")
//			},
//		}
//
//		// use mockedAddTaskService in code that requires AddTaskService
//		// and then make assertions.
//
//	}
type AddTaskServiceMock struct {
	// AddTaskFunc mocks the AddTask method.
	AddTaskFunc func(ctx context.Context, title string) (*entity.Task, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddTask holds details about calls to the AddTask method.
		AddTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Title is the title argument value.
			Title string
		}
	}
	lockAddTask sync.RWMutex
}

// AddTask calls AddTaskFunc.
func (mock *AddTaskServiceMock) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	if mock.AddTaskFunc == nil {
		panic("AddTaskServiceMock.AddTaskFunc: method is nil but AddTaskService.AddTask was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Title string
	}{
		Ctx:   ctx,
		Title: title,
	}
	mock.lockAddTask.Lock()
	mock.calls.AddTask = append(mock.calls.AddTask, callInfo)
	mock.lockAddTask.Unlock()
	return mock.AddTaskFunc(ctx, title)
}

// AddTaskCalls gets all the calls that were made to AddTask.
// Check the length with:
//
//	len(mockedAddTaskService.AddTaskCalls())
func (mock *AddTaskServiceMock) AddTaskCalls() []struct {
	Ctx   context.Context
	Title string
} {
	var calls []struct {
		Ctx   context.Context
		Title string
	}
	mock.lockAddTask.RLock()
	calls = mock.calls.AddTask
	mock.lockAddTask.RUnlock()
	return calls
}
