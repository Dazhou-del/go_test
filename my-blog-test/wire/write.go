package wire

import (
	"github.com/google/wire"
	"test/my-blog-test/internal/post"
)

func InitializeApp() error {
	wire.Build(
		post.NewGinEngine,

	)
}
