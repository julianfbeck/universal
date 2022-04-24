import {
  AppShell,
  Navbar,
  Header,
  Group,
  ActionIcon,
  MediaQuery,
  Burger,
  useMantineTheme,
  Text,
} from "@mantine/core";
import { Outlet } from "@remix-run/react";
import { useState } from "react";
import { Users } from "tabler-icons-react";
import { MainLinks } from "./MainLinks";

export default function AppNavbar() {
  const [opened, setOpened] = useState(false);
  const theme = useMantineTheme();

  return (
    <AppShell
      padding="md"
      navbarOffsetBreakpoint="sm"
      asideOffsetBreakpoint="sm"
      fixed
      navbar={
        <Navbar
          p="md"
          hiddenBreakpoint="sm"
          hidden={!opened}
          width={{ sm: 200, lg: 250 }}
        >
          <Navbar.Section grow mt="xs">
            <MainLinks />
          </Navbar.Section>
        </Navbar>
      }
      header={
        <Header height={70} p="md">
          <div
            style={{ display: "flex", alignItems: "center", height: "100%" }}
          >
            <MediaQuery largerThan="sm" styles={{ display: "none" }}>
              <Burger
                opened={opened}
                onClick={() => setOpened((o) => !o)}
                size="sm"
                color={theme.colors.gray[6]}
                mr="xl"
              />
            </MediaQuery>

            <Group sx={{ height: "100%" }} px={20} position="apart">
              <ActionIcon>
                <Users />
              </ActionIcon>
              <Text
                sx={(theme) => ({
                  color:
                    theme.colorScheme === "dark"
                      ? theme.colors.dark[0]
                      : theme.black,
                })}
              >
                Beju App
              </Text>
            </Group>
          </div>
        </Header>
      }
      styles={(theme) => ({
        main: {
          backgroundColor:
            theme.colorScheme === "dark"
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
          color:
            theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,
        },
      })}
    >
      <Outlet />
    </AppShell>
  );
}
