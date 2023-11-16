import { OnboardingPage } from "../page";

import { Meta } from "@storybook/react";
import { Theme } from "@/styles/__stories__/Theme";

const meta = {
  title: "Pages/Onboarding",
  parameters: {
    layout: "fullscreen",
  },
} satisfies Meta;

export default meta;

export const Overview = () => (
  <Theme.Light>
    <OnboardingPage />
  </Theme.Light>
);
