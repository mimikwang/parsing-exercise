from abc import ABC, abstractmethod

from pydantic import BaseModel

from parse import output


class Parse(ABC):
    """Abstract base class for parser formats"""

    @abstractmethod
    def parse(self) -> output.Output:
        """Required method for class"""
        raise NotImplementedError


class FormatA(Parse, BaseModel):
    id: int
    name: str
    country: str
    age: int
    weight: float
    height: float

    def parse(self) -> output.Output:
        """Convert Format A to Output"""
        bmi = self.weight / self.height**2
        return output.Output(
            id=self.id,
            name=self.name,
            age=self.age,
            bmi=output.Bmi.from_bmi(bmi)
        )


class FormatBInfo(BaseModel):
    name: str
    country: str


class FormatBData(BaseModel):
    bmi: float
    age: int


class FormatB(Parse, BaseModel):
    id: int
    info: FormatBInfo
    data: FormatBData

    def parse(self) -> output.Output:
        """Convert FormatB to Output"""
        return output.Output(
            id=self.id,
            name=self.info.name,
            age=self.data.age,
            bmi=output.Bmi.from_bmi(self.data.bmi)
        )
