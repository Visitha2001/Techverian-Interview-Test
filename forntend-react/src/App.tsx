import ItemForm from "@/components/ItemForm"
import ItemList from "@/components/ItemList"
import { useState } from "react"

function App() {
  const [refreshKey, setRefreshKey] = useState(0)

  const handleAdded = () => setRefreshKey((k) => k + 1)

  return (
    <>
      <div className='flex flex-col justify-center items-center sm:h-screen h-auto min-h-screen sm:py-0 py-10 sm:px-0 px-6 sm:w-7xl w-full mx-auto'>
        <h1 className='text-3xl font-bold'>Add Items For Database And Get Stats</h1>
        <div className="flex w-full flex-col gap-4 mt-10 sm:flex-row">
          <ItemForm onAdded={handleAdded} />
          <ItemList refreshKey={refreshKey} />
        </div>
      </div>
    </>
  )
}

export default App
