import React from "react";
import { Group, ActionIcon, useMantineColorScheme, Box } from "@mantine/core";
import { Sun, MoonStars } from "tabler-icons-react";

export function Brand() {
  return (
    <Box
      sx={(theme) => ({
        paddingLeft: theme.spacing.xs,
        paddingRight: theme.spacing.xs,
        paddingBottom: theme.spacing.lg,
        borderBottom: `1px solid ${
          theme.colorScheme === "dark"
            ? theme.colors.dark[4]
            : theme.colors.gray[2]
        }`,
      })}
    >
      <Group position="apart">Bejuuuuu</Group>
    </Box>
  );
}
