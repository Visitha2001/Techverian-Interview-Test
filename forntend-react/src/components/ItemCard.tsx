import type { Item } from "@/types"

const ItemCard = ({item}: {item: Item}) => {
  return (
    <div className='w-full bg-gray-700 rounded-lg p-4'>
        <div className="flex flex-row justify-between gap-2">
            <h2 className="text-lg font-semibold">{item.name}</h2>
            <p className="text-gray-300">{item.price}</p>
        </div>
    </div>
  )
}

export default ItemCard