import React from "react";
import {
  GitPullRequest,
  AlertCircle,
  Messages,
  Database,
  ShoppingCart,
  ShoppingCartPlus,
  Settings,
  List,
} from "tabler-icons-react";
import { ThemeIcon, UnstyledButton, Group, Text } from "@mantine/core";
import { Link, useLocation } from "@remix-run/react";

interface MainLinkProps {
  icon: React.ReactNode;
  color: string;
  label: string;
  href: string;
}

function MainLink({ icon, color, label, href }: MainLinkProps) {
  const location = useLocation();
  return (
    <Link to={href} style={{ textDecoration: "none" }}>
      <UnstyledButton
        sx={(theme) => ({
          display: "block",
          width: "100%",
          padding: theme.spacing.xs,
          borderRadius: theme.radius.sm,
          color:
            theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,

          "&:hover": {
            backgroundColor:
              theme.colorScheme === "dark"
                ? theme.colors.dark[6]
                : theme.colors.gray[0],
          },
          backgroundColor:
            location.pathname === href
              ? theme.colors.dark[6]
              : "transparent",
        })}
      >
        <Group>
          <ThemeIcon color={color} variant="light">
            {icon}
          </ThemeIcon>
          <Text size="sm">{label}</Text>
        </Group>
      </UnstyledButton>
    </Link>
  );
}

const data = [
  {
    icon: <ShoppingCart size={16} />,
    color: "blue",
    label: "Overview",
    href: "/",
  },
  {
    icon: <ShoppingCartPlus size={16} />,
    color: "teal",
    label: "Add Item",
    href: "/add",
  },
  {
    icon: <List size={16} />,
    color: "violet",
    label: "Wishlist",
    href: "/wishlist",
  },
  {
    icon: <Settings size={16} />,
    color: "grape",
    label: "Settings",
    href: "/settings",
  },
];

export function MainLinks() {
  const links = data.map((link) => <MainLink {...link} key={link.label} />);
  return <div>{links}</div>;
}
