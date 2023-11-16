import { OnboardingPage } from "../page";

import { Meta } from "@storybook/react";
import { Theme } from "@/styles/__stories__/Theme";

const meta = {
  title: "Pages/Onboarding",
} satisfies Meta;

export default meta;

export const Light = () => (
  <Theme.Light full>
    <OnboardingPage />
  </Theme.Light>
);

export const Dark = () => (
  <Theme.Dark full>
    <OnboardingPage />
  </Theme.Dark>
);
