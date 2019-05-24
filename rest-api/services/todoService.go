package services

import (
	"rest-api/models"
	"rest-api/utils"
	"sync"
)

func ListTodo(userID, offset, limit int) ([]*models.TodoInfo, uint64, error) {
	infos := make([]*models.TodoInfo, 0)
	todos, count, err := models.ListTodo(userID, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, todo := range todos {
		ids = append(ids, todo.ID)
	}

	wg := sync.WaitGroup{}
	todoList := models.TodoList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*models.TodoInfo, len(todos)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, td := range todos {
		wg.Add(1)
		go func(td *models.TodoModel) {
			defer wg.Done()

			_, err := utils.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			todoList.Lock.Lock()
			defer todoList.Lock.Unlock()
			todoList.IdMap[td.ID] = &models.TodoInfo{
				ID:        td.ID,
				Title:     td.Title,
				Status:    td.Status,
				UserID:    td.UserID,
				CreatedAt: td.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: td.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(td)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, todoList.IdMap[id])
	}

	return infos, count, nil
}
