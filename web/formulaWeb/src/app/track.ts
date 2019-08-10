import { Deserializable } from "./deserializable.model";

export class Track {
  id: number;
  name: string;
  length: number;
  country: string;
  official_record: string;

  deserialize(input: any) {
    Object.assign(this, input);
    return this;
  }
}
