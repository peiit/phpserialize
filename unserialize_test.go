package phpserialize_test

//func decodeNull(input []byte, expectedError error) {
//	var result interface{}
//	err := phpserialize.Unmarshal(input, &result)
//
//	Expect(result).To(BeNil())
//
//	err = phpserialize.UnmarshalNil(input)
//	if expectedError == nil {
//		Expect(err).ToNot(HaveOccurred())
//	} else {
//		Expect(err).To(HaveOccurred())
//		Expect(err.Error()).To(Equal(err.Error()))
//	}
//}
//
//func decodeBool(input []byte, output bool, expectedError error) {
//	var result bool
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		Expect(err).ToNot(HaveOccurred())
//		Expect(result).To(Equal(output))
//	} else {
//		Expect(err).To(HaveOccurred())
//		Expect(err.Error()).To(Equal(expectedError.Error()))
//	}
//}
//
//func decodeInt(input []byte, output int, expectedError error) {
//	{
//		var result int
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(output))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result int8
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(int8(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result int16
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(int16(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result int32
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(int32(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result int64
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(int64(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result uint8
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(uint8(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result uint16
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(uint16(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result uint32
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(uint32(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result uint64
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(uint64(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//}
//
//func decodeFloat(input []byte, output float64, expectedError error) {
//	{
//		var result float32
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(float32(output)))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//
//	{
//		var result float64
//		err := phpserialize.Unmarshal(input, &result)
//
//		if expectedError == nil {
//			Expect(err).ToNot(HaveOccurred())
//			Expect(result).To(Equal(output))
//		} else {
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(Equal(expectedError.Error()))
//		}
//	}
//}
//
//func decodeString(input []byte, output string, expectedError error) {
//	var result string
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		Expect(err).ToNot(HaveOccurred())
//		Expect(result).To(Equal(output))
//	} else {
//		Expect(err).To(HaveOccurred())
//		Expect(err.Error()).To(Equal(expectedError.Error()))
//	}
//}
//
//func decodeBinary(input []byte, output []byte, expectedError error) {
//	var result []byte
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		Expect(err).ToNot(HaveOccurred())
//		Expect(result).To(Equal(output))
//	} else {
//		Expect(err).To(HaveOccurred())
//		Expect(err.Error()).To(Equal(expectedError.Error()))
//	}
//}
//
//func decodeArray(input []byte, output []interface{}, expectedError error) {
//	var result []interface{}
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		Expect(err).ToNot(HaveOccurred())
//		Expect(len(result)).To(Equal(len(output)))
//		for k, _ := range result {
//			// Ginkgo has a safety feature when comparing two nil
//			// values for equality. You have to use the nil
//			// assertions. I this case it's not important.
//			if result[k] != nil || output[k] != nil {
//				Expect(result[k]).To(BeEquivalentTo(output[k]))
//			}
//		}
//	} else {
//		Expect(err).To(HaveOccurred())
//		Expect(err.Error()).To(Equal(expectedError.Error()))
//	}
//}
//
//func decodeAssociativeArray(input []byte, output map[interface{}]interface{}, expectedError error) {
//	result := make(map[interface{}]interface{})
//	err := phpserialize.Unmarshal(input, &result)
//
//	if expectedError == nil {
//		Expect(err).ToNot(HaveOccurred())
//		Expect(reflect.DeepEqual(result, output)).To(BeTrue())
//	} else {
//		Expect(err).To(HaveOccurred())
//		Expect(err.Error()).To(Equal(expectedError.Error()))
//	}
//}
//
//var _ = Describe("phpserialize", func() {
//	Describe("Unmarshal - unserialize()", func() {
//		DescribeTable("decode null",
//			decodeNull,
//
//			Entry("null", []byte("N;"), nil),
//
//			Entry("not null", []byte("b:0;"), errors.New("not null")),
//		)
//
//		DescribeTable("decode bool",
//			decodeBool,
//
//			Entry("true", []byte("b:1;"), true, nil),
//			Entry("false", []byte("b:0;"), false, nil),
//
//			Entry("not a boolean", []byte("N;"), true, errors.New("not a boolean")),
//		)
//
//		DescribeTable("decode int",
//			decodeInt,
//
//			Entry("0", []byte("i:0;"), 0, nil),
//			Entry("5", []byte("i:5;"), 5, nil),
//			Entry("-8", []byte("i:-8;"), -8, nil),
//			Entry("1000000", []byte("i:1000000;"), 1000000, nil),
//
//			Entry("not an integer", []byte("N;"), 0, errors.New("not an integer")),
//		)
//
//		DescribeTable("decode float",
//			decodeFloat,
//
//			Entry("3.2", []byte("d:3.2;"), 3.2, nil),
//			Entry("10.0", []byte("d:10;"), 10.0, nil),
//			Entry("123.456789", []byte("d:123.456789;"), 123.456789, nil),
//			Entry("1.23e9", []byte("d:1230000000;"), 1.23e9, nil),
//			Entry("-17.23", []byte("d:3.2;"), 3.2, nil),
//
//			Entry("not a float", []byte("N;"), 0.0, errors.New("not a float")),
//		)
//
//		DescribeTable("decode string",
//			decodeString,
//
//			Entry("''", []byte("s:0:\"\";"), "", nil),
//			Entry("'Hello world'", []byte("s:11:\"Hello world\";"),
//				"Hello world", nil),
//			Entry("'Björk Guðmundsdóttir'",
//				[]byte("s:23:\"Bj\\xc3\\xb6rk Gu\\xc3\\xb0mundsd\\xc3\\xb3ttir\";"),
//				"Björk Guðmundsdóttir", nil),
//
//			Entry("not a string", []byte("N;"), "", errors.New("not a string")),
//		)
//
//		DescribeTable("decode binary",
//			decodeBinary,
//
//			Entry("[]byte: \\001\\002\\003", []byte("s:3:\"\\x01\\x02\\x03\";"),
//				[]byte{1, 2, 3}, nil),
//
//			Entry("not a string", []byte("N;"), []byte{}, errors.New("not a string")),
//		)
//
//		DescribeTable("decode array (slice)",
//			decodeArray,
//
//			Entry("[]interface{}: [7.89]", []byte("a:1:{i:0;d:7.89;}"),
//				[]interface{}{7.89}, nil),
//			Entry("[]interface{}: [7, 8, 9]",
//				[]byte("a:3:{i:0;i:7;i:1;i:8;i:2;i:9;}"),
//				[]interface{}{7, 8, 9}, nil),
//			Entry("[]interface{}: [7.2, 'foo']",
//				[]byte("a:2:{i:0;d:7.2;i:1;s:3:\"foo\";}"),
//				[]interface{}{7.2, "foo"}, nil),
//			Entry("[]interface{}: [null]",
//				[]byte("a:1:{i:0;N;}"),
//				[]interface{}{nil}, nil),
//			Entry("[]interface{}: [true, false]",
//				[]byte("a:2:{i:0;b:1;i:1;b:0;}"),
//				[]interface{}{true, false}, nil),
//
//			Entry("cannot decode map as slice",
//				[]byte("a:2:{i:0;b:1;i:5;b:0;}"),
//				[]interface{}{},
//				errors.New("cannot decode map as slice")),
//			Entry("not an array", []byte("N;"), []interface{}{},
//				errors.New("not an array")),
//		)
//
//		DescribeTable("decode associative array (map)",
//			decodeAssociativeArray,
//
//			Entry("map[interface{}]interface{}: {'foo': 10, 'bar': 20}",
//				[]byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}"),
//				map[interface{}]interface{}{"foo": int64(10), "bar": int64(20)},
//				nil),
//			Entry("map[interface{}]interface{}: {1: 10, 2: 'foo'}",
//				[]byte("a:2:{i:1;i:10;i:2;s:3:\"foo\";}"),
//				map[interface{}]interface{}{int64(1): int64(10), int64(2): "foo"},
//				nil),
//
//			Entry("not an array", []byte("N;"),
//				map[interface{}]interface{}{},
//				errors.New("not an array")),
//		)
//
//		Describe("decode object", func() {
//			It("struct1{Foo int, Bar Struct2{Qux float64}, hidden bool}", func() {
//				data := "O:7:\"struct1\":3:{s:3:\"foo\";i:10;s:3:\"bar\";O:7:\"Struct2\":1:{s:3:\"qux\";d:1.23;}s:3:\"baz\";s:3:\"yay\";}"
//				var result struct1
//				err := phpserialize.Unmarshal([]byte(data), &result)
//				Expect(err).ToNot(HaveOccurred())
//
//				Expect(result.Foo).To(Equal(10))
//				Expect(result.Bar.Qux).To(Equal(1.23))
//				Expect(result.Baz).To(Equal("yay"))
//			})
//		})
//	})
//})