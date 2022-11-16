import pytest

from parse import output


class TestBmi:
    @pytest.mark.parametrize("bmi,expected", [
        (10, output.Bmi.UNDERWEIGHT),
        (23, output.Bmi.NORMAL),
        (28, output.Bmi.OVERWEIGHT),
        (600, output.Bmi.VERY_OVERWEIGHT),
    ])
    def test_from_bmi(self, bmi, expected):
        assert output.Bmi.from_bmi(bmi) == expected
