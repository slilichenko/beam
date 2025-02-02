// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipelinex

import (
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/reflectx"
	pipepb "github.com/apache/beam/sdks/v2/go/pkg/beam/model/pipeline_v1"
)

func shallowClonePipeline(p *pipepb.Pipeline) *pipepb.Pipeline {
	ret := pipepb.Pipeline_builder{
		Components:   shallowCloneComponents(p.GetComponents()),
		Requirements: reflectx.ShallowClone(p.GetRequirements()).([]string),
	}.Build()
	ret.RootTransformIds, _ = reflectx.ShallowClone(p.GetRootTransformIds()).([]string)
	return ret
}

func shallowCloneComponents(comp *pipepb.Components) *pipepb.Components {
	ret := &pipepb.Components{}
	ret.Transforms, _ = reflectx.ShallowClone(comp.GetTransforms()).(map[string]*pipepb.PTransform)
	ret.Pcollections, _ = reflectx.ShallowClone(comp.GetPcollections()).(map[string]*pipepb.PCollection)
	ret.WindowingStrategies, _ = reflectx.ShallowClone(comp.GetWindowingStrategies()).(map[string]*pipepb.WindowingStrategy)
	ret.Coders, _ = reflectx.ShallowClone(comp.GetCoders()).(map[string]*pipepb.Coder)
	ret.Environments, _ = reflectx.ShallowClone(comp.GetEnvironments()).(map[string]*pipepb.Environment)
	return ret
}

// ShallowClonePTransform makes a shallow copy of the given PTransform.
func ShallowClonePTransform(t *pipepb.PTransform) *pipepb.PTransform {
	if t == nil {
		return nil
	}

	ret := pipepb.PTransform_builder{
		UniqueName:  t.GetUniqueName(),
		Spec:        t.GetSpec(),
		DisplayData: t.GetDisplayData(),
		Annotations: t.GetAnnotations(),
	}.Build()
	ret.Subtransforms, _ = reflectx.ShallowClone(t.GetSubtransforms()).([]string)
	ret.Inputs, _ = reflectx.ShallowClone(t.GetInputs()).(map[string]string)
	ret.Outputs, _ = reflectx.ShallowClone(t.GetOutputs()).(map[string]string)
	ret.SetEnvironmentId(t.GetEnvironmentId())
	return ret
}
