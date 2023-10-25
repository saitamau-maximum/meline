/**
 * Utility function for conditionally joining CSS class names together.
 * @param classes - Array of CSS class names or objects with class names as keys and boolean values as values.
 * @returns Joined CSS class names.
 */
export function clsx(
  ...classes: Array<string | Record<string, boolean> | undefined | null>
): string {
  const classNames: string[] = [];

  for (const cls of classes) {
    if (typeof cls === "string") {
      classNames.push(cls);
    } else if (cls && typeof cls === "object") {
      classNames.push(
        ...Object.entries(cls)
          .filter(([_, value]) => value)
          .map(([key]) => key)
      );
    }
  }

  return classNames.join(" ");
}
