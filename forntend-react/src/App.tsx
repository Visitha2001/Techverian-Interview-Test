import ItemForm from "@/components/ItemForm"
import ItemList from "@/components/ItemList"
import { useState } from "react"

function App() {
  const [refreshKey, setRefreshKey] = useState(0)

  const handleAdded = () => setRefreshKey((k) => k + 1)

  return (
    <>
      <div className='flex flex-col justify-center items-center h-screen w-7xl mx-auto'>
        <h1 className='text-3xl font-bold'>Add Items For Database And Get Stats</h1>
        <div className="flex w-full flex-col-2 gap-4 mt-10">
          <ItemForm onAdded={handleAdded} />
          <ItemList refreshKey={refreshKey} />
        </div>
      </div>
    </>
  )
}

export default App
