// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"log"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
