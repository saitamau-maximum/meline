import { Meta } from "@storybook/react";
import { Textarea } from "..";

const meta = {
  title: "Components/UI/Textarea",
} satisfies Meta;

export default meta;

export const Overview = () => (
  <Textarea
    id="textarea"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
  />
);

export const WithError = () => (
  <Textarea
    id="textarea"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
    error="お名前は必須です。"
  />
);

export const WithDisabled = () => (
  <Textarea
    id="textarea"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
    disabled
  />
);

export const WithRequired = () => (
  <Textarea
    id="textarea"
    placeholder="山田 太郎"
    label="お名前"
    description="お名前を漢字フルネームでお書きください。苗字と名前の間にはスペースが必要です。"
    required
  />
);
