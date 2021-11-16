// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package breaking

import (
	"github.com/Berrserker/prototool/internal/extract"
	"github.com/Berrserker/prototool/internal/text"
)

func checkMessageFieldsSameLabel(addFailure func(*text.Failure), from *extract.PackageSet, to *extract.PackageSet) error {
	return forEachMessageFieldPair(addFailure, from, to, checkMessageFieldsSameLabelMessageField)
}

func checkMessageFieldsSameLabelMessageField(addFailure func(*text.Failure), from *extract.MessageField, to *extract.MessageField) error {
	fromLabel := from.ProtoMessage().Label
	toLabel := to.ProtoMessage().Label
	if fromLabel != toLabel {
		fromLabelString, err := getMessageFieldLabelString(fromLabel)
		if err != nil {
			return err
		}
		toLabelString, err := getMessageFieldLabelString(toLabel)
		if err != nil {
			return err
		}
		addFailure(newMessageFieldsSameLabelFailure(from.Message().FullyQualifiedName(), from.ProtoMessage().Number, fromLabelString, toLabelString))
		return nil
	}
	return nil
}

func newMessageFieldsSameLabelFailure(messageName string, fieldNumber int32, fromLabelString string, toLabelString string) *text.Failure {
	return newTextFailuref(`Message field "%d" on message %q changed from %q to %q.`, fieldNumber, messageName, fromLabelString, toLabelString)
}
