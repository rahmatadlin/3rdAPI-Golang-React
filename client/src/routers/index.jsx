import React from "react";
import { createBrowserRouter } from "react-router-dom";
import BaseLayout from "../views/BaseLayout";
import HomePage from "../views/HomePage";

const SERVER_URL = "http://localhost:4000";

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage url={SERVER_URL} />,
  },
  {
    path: "/books",
    element: <HomePage url={SERVER_URL} />,
  },
  {
    element: <BaseLayout />,
    children: [
      {
        path: "/books",
        element: <HomePage url={SERVER_URL} />,
      },
    ],
  },
]);

export default router;
