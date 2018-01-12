/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package diff

import "reflect"

func (cl *Changelog) diffStruct(path []string, a, b reflect.Value) error {
	if a.Kind() != b.Kind() {
		return ErrTypeMismatch
	}

	for i := 0; i < a.NumField(); i++ {
		field := a.Type().Field(i)
		tname := tagName(field)

		if tname == "-" || hasTagOption(field, "immutable") {
			continue
		}

		af := a.Field(i)
		bf := b.FieldByName(field.Name)

		fpath := append(path, tname)

		err := cl.diff(fpath, af, bf)
		if err != nil {
			return err
		}
	}

	return nil
}
