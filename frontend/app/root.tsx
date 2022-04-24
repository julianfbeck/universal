import type { LinksFunction, MetaFunction } from "@remix-run/node";
import { MantineProvider, ColorSchemeProvider } from "@mantine/core";

import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";
import AppNavbar from "./components/Navbar";
import styles from "~/styles/global.css";

export const links: LinksFunction = () => {
  return [
    {
      rel: "stylesheet",
      href: styles,
    },
  ];
};
export default function App() {
  return (
    <MantineProvider theme={{ colorScheme: "dark" }}>
      <html lang="en">
        <head>
          <Meta />
          <Links />
        </head>
        <body>
          <AppNavbar />
          <ScrollRestoration />
          <Scripts />
          <LiveReload />
        </body>
      </html>
    </MantineProvider>
  );
}
