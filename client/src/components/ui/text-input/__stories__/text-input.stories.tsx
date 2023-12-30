import { Meta } from "@storybook/react";
import { TextInput } from "..";

const meta = {
  title: "Components/UI/TextInput",
} satisfies Meta;

export default meta;

export const Overview = () => (
  <TextInput
    id="text-input"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
  />
);

export const WithError = () => (
  <TextInput
    id="text-input"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
    error="お名前は必須です。"
  />
);

export const WithDisabled = () => (
  <TextInput
    id="text-input"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
    disabled
  />
);

export const WithRequired = () => (
  <TextInput
    id="text-input"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
    required
  />
);
