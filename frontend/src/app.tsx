import { CreateRoom } from "./pages/create-room"
import { Room } from "./pages/room"
import { queryClient } from "./lib/react-query"

import { createBrowserRouter, RouterProvider } from "react-router-dom"
import { Toaster } from "sonner"
import { QueryClientProvider } from "@tanstack/react-query"

const router = createBrowserRouter([
  {
    path: '/',
    element: <CreateRoom />
  },
  {
    path: '/room/:roomId',
    element: <Room />
  }
])

export function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      <Toaster
        toastOptions={{
          unstyled: true,
          classNames: {
            toast: "font-medium w-[356px] min-h-[50px] bg-zinc-800 text-zinc-300 gap-1.5 text-white px-4 py-3 rounded-md shadow-lg flex items-center ml-auto",
            title: "text-[13px]",
            error: "text-red-400",
            success: "text-green-400",
            warning: "text-yellow-400",
            info: "text-blue-400",
            actionButton: "bg-orange-400 hover:bg-orange-500 font-bold text-[12px] text-orange-950 px-2 py-1 rounded-md",
          },
        }}
      />
    </QueryClientProvider>
  )
}
