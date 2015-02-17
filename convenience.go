package mbapi

import "fmt"

func (b *Backend) Movies() ([]BaseItemDto, error) {
	if b.auth == nil {
		return nil, fmt.Errorf("Not logged in")
	}
	query := &ItemQuery{
		UserId:           b.auth.User.Id,
		Recursive:        true,
		SortBy:           []string{"SortName"},
		Fields:           []string{"MediaSources"},
		IncludeItemTypes: []string{"Movie"},
	}
	resp, err := b.NewRequest(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Movies: %v", err)
	}

	var items ItemsResult

	err = RenderResponse(resp, &items)

	return items.Items, err
}

func (b *Backend) TVShows() ([]BaseItemDto, error) {
	if b.auth == nil {
		return nil, fmt.Errorf("Not logged in")
	}
	query := &ItemQuery{
		UserId:           b.auth.User.Id,
		Recursive:        true,
		SortBy:           []string{"SortName"},
		Fields:           []string{"MediaSources"},
		IncludeItemTypes: []string{"Series"},
	}
	resp, err := b.NewRequest(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Movies: %v", err)
	}

	var items ItemsResult

	err = RenderResponse(resp, &items)

	return items.Items, err
}

func (b *Backend) ShowById(id string) (*BaseItemDto, error) {
	if b.auth == nil {
		return nil, fmt.Errorf("Not logged in")
	}
	query := &ItemQuery{
		UserId:           b.auth.User.Id,
		Recursive:        true,
		IncludeItemTypes: []string{"Series"},
		Ids:              []string{id},
	}
	resp, err := b.NewRequest(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get TV-Show: %v", err)
	}
	var items ItemsResult
	err = RenderResponse(resp, &items)
	if len(items.Items) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &items.Items[0], nil

}

func (b *Backend) SeasonById(id string) (*BaseItemDto, error) {
	if b.auth == nil {
		return nil, fmt.Errorf("Not logged in")
	}
	query := &ItemQuery{
		UserId:           b.auth.User.Id,
		Recursive:        true,
		IncludeItemTypes: []string{"Season"},
		Ids:              []string{id},
	}
	resp, err := b.NewRequest(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Season: %v", err)
	}
	var items ItemsResult
	err = RenderResponse(resp, &items)
	if len(items.Items) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &items.Items[0], nil

}

func (b *Backend) EpisodeById(id string) (*BaseItemDto, error) {
	if b.auth == nil {
		return nil, fmt.Errorf("Not logged in")
	}
	query := &ItemQuery{
		UserId:           b.auth.User.Id,
		Recursive:        true,
		IncludeItemTypes: []string{"Episode"},
		Ids:              []string{id},
		Fields:           []string{"MediaSources"},
	}
	resp, err := b.NewRequest(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Episode: %v", err)
	}
	var items ItemsResult
	err = RenderResponse(resp, &items)
	if len(items.Items) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &items.Items[0], nil

}

func (b *Backend) GetUsers() ([]UserDto, error) {
	r, err := b.NewRequest(&ListUsersQuery{})
	if err != nil {
		return nil, fmt.Errorf("Failed to get Users: %v", err)
	}
	var users []UserDto
	err = RenderResponse(r, &users)

	return users, err
}
