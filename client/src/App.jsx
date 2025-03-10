import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import './App.css'

export const BASE_URL = import.meta.env.VITE_ENV == "production" ? "/api" : import.meta.env.VITE_API_URL + ":" + import.meta.env.VITE_API_PORT + "/api"

function App() {
  const queryClient = useQueryClient()

  const { data: resume, isLoading } = useQuery({
    queryKey: ["resume"],
    queryFn: async () => {
      try {
        const res = await fetch(BASE_URL + `/resume`)
        const data = await res.json()

        if (!res.ok) {
          throw new Error(data.error || "Something went wrong")
        }

        return data || []
      } catch (error) {
        console.log(error)
      }
    }
  })

  const { mutate: updateResume, isPending: isUpdating } = useMutation({
    mutationKey: ["updateResume"],
    mutationFn: async () => {
      try {
        const res = await fetch(`${BASE_URL}/resume/update`, {
          method: "PUT"
        })

        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something went wrong")
        }
        return data || [];
      } catch (error) {
        console.log(error)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["resume"] })
    }
  })


  return (
    <>
      <p>VK: {resume?.vk_stats}</p>
      <p>GITHUB: {resume?.github_stats}</p>
      <p>CODEFORCES: {resume?.codeforces_stats}</p>
      <p>UPDATED AT: {resume?.updated_at}</p>
      <button onClick={updateResume}>{!isUpdating ? "Update" : "Updating.." }</button>
    </>
  )
}

export default App
