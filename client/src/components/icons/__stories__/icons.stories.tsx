import { Meta } from "@storybook/react";
import { GithubIcon } from "../github";
import { MaximumIcon } from "../maximum";
import { vars } from "@/styles";

const meta = {
  title: "Components/Icons",
} satisfies Meta;

export default meta;

const ICONS = {
  GithubIcon,
  MaximumIcon,
};

export const Overview = () => (
  <div
    style={{
      display: "flex",
      alignItems: "center",
      flexWrap: "wrap",
      width: "100%",
      gap: vars.spacing[8],
    }}
  >
    {Object.entries(ICONS).map(([name, Icon]) => (
      <div
        key={name}
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: vars.spacing[4],
        }}
      >
        <span
          style={{
            fontSize: vars.font.size["2xl"],
          }}
        >
          <Icon />
        </span>
        <span>{name}</span>
      </div>
    ))}
  </div>
);
