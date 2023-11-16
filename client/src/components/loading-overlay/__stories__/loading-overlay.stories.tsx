import { Meta } from "@storybook/react";
import { Theme } from "@/styles/__stories__/Theme";
import { LoadingOverlay } from "..";
import { LoadingOverlayContext } from "@/providers/loading-overlay";

const meta = {
  title: "Components/LoadingOverlay",
} satisfies Meta;

export default meta;

const Common = () => (
  <LoadingOverlayContext.Provider
    value={{ isLoading: true, setIsLoading: () => {} }}
  >
    <LoadingOverlay />
  </LoadingOverlayContext.Provider>
);

export const Light = () => (
  <Theme.Light>
    <Common />
  </Theme.Light>
);

export const Dark = () => (
  <Theme.Dark>
    <Common />
  </Theme.Dark>
);
