package post

type IPostServiceInterface interface {
	GetPostById(id string) (string, error)
}

type IPostService struct{}

func (p *IPostService) GetPostById(id string) (string, error) {
	return "文章内容", nil
}
