package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseCopyrightHolder(t *testing.T) {
	e := copyrightRegexHeuristic{}
	assert.Equal(t, "2016 gotohr", e.parseCopyrightHolder("Copyright (c) 2016 gotohr"))
	assert.Equal(t, "John Doe <john.d@example.com>", e.parseCopyrightHolder("Copyright (c) John Doe <john.d@example.com>"))
	assert.Equal(t, "2009 The Go Authors. All rights reserved.", e.parseCopyrightHolder("Copyright (c) 2009 The Go Authors. All rights reserved."))
	assert.Equal(t, "2011 John Doe <john.d@example.com>", e.parseCopyrightHolder("Copyright © 2011 John Doe <john.d@example.com>"))
}
