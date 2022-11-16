import pytest

from parse import input, output


class TestFormatA:
    def test_parse_raw(self):
        raw = '{"id": 23, "name": "foo", "country": "bar", "age": 42, "weight": 23.1, "height": 10.1}'
        actual = input.FormatA.parse_raw(raw)
        expected = input.FormatA(
            id=23, 
            name="foo", 
            country="bar", 
            age=42, 
            weight=23.1, 
            height=10.1
        )

        assert actual == expected

    def test_parse_raw_missing_fields(self):
        raw = '{"id": 23, "name": "foo", "country": "bar", "age": 42, "weight": 23.1}'
        with pytest.raises(Exception):
            input.FormatA.parse_raw(raw)

    def test_parse(self):
        formatA = input.FormatA(
            id=23, 
            name="foo", 
            country="bar", 
            age=42, 
            weight=23.1, 
            height=10.1
        )
        actual = formatA.parse()
        expected = output.Output(
            id=23, 
            name="foo", 
            age=42, 
            bmi=output.Bmi.UNDERWEIGHT
        )

        assert actual == expected


class TestFormatB:
    def test_parse_raw(self):
        raw = '{"id": 23, "info": {"name": "foo", "country": "bar"}, "data": {"bmi": 3.1, "age": 42} }'
        actual = input.FormatB.parse_raw(raw)
        expected = input.FormatB(
            id=23, 
            info=input.FormatBInfo(name="foo", country="bar"), 
            data=input.FormatBData(bmi=3.1, age=42)
        )

        assert actual == expected
    
    def test_parse_raw_missing_fields(self):
        raw = '{"info": {"name": "foo", "country": "bar"}, "data": {"bmi": 3.1, "age": 42} }'
        with pytest.raises(Exception):
            input.FormatB.parse_raw(raw)

    def test_parse(self):
        formatB = input.FormatB(
            id=23,
            info=input.FormatBInfo(name="foo", country="bar"),
            data=input.FormatBData(bmi=3.1, age=42)
        )
        actual = formatB.parse()
        expected = output.Output(
            id=23,
            name="foo",
            age=42,
            bmi=output.Bmi.UNDERWEIGHT
        )

        assert actual == expected
