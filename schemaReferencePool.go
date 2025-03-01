// Copyright 2015 xeipuuv ( https://github.com/xeipuuv )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author           xeipuuv
// author-github    https://github.com/xeipuuv
// author-mail      xeipuuv@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Pool of referenced schemas.
//
// created          25-06-2013

package gojsonschema

import (
	"fmt"
)

type SchemaReferencePool struct {
	documents map[string]*SubSchema
}

func newSchemaReferencePool() *SchemaReferencePool {

	p := &SchemaReferencePool{}
	p.documents = make(map[string]*SubSchema)

	return p
}

func (p *SchemaReferencePool) Get(ref string) (r *SubSchema, o bool) {

	if internalLogEnabled {
		internalLog(fmt.Sprintf("Schema Reference ( %s )", ref))
	}

	if sch, ok := p.documents[ref]; ok {
		if internalLogEnabled {
			internalLog(fmt.Sprintf(" From Pool"))
		}
		return sch, true
	}

	return nil, false
}

func (p *SchemaReferencePool) Add(ref string, sch *SubSchema) {

	if internalLogEnabled {
		internalLog(fmt.Sprintf("Add Schema Reference %s to Pool", ref))
	}
	if _, ok := p.documents[ref]; !ok {
		p.documents[ref] = sch
	}
}
