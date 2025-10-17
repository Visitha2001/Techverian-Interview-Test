import { itemService } from "@/services/itemService";
import ItemCard from "./ItemCard";
import { useEffect, useState } from "react";
import type { Item } from "@/types";
import { Spinner } from "./ui/spinner";

type ItemListProps = { refreshKey?: number }

const ItemList = ({ refreshKey }: ItemListProps) => {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const [averageCost, setAverageCost] = useState(0);
  const [totalCost, setTotalCost] = useState(0);

  useEffect(() => {
    let isMounted = true
    const fetchItems = async () => {
      try {
        const response = await itemService.getAllItems();
        if (isMounted) setItems(response.items);
      } catch (error) {
        console.error("Failed to fetch items:", error);
        if (isMounted) setItems([]);
      } finally {
        if (isMounted) setLoading(false);
      }
    };
    
    const getSummary = async () => {
      const response = await itemService.getSummary();
      if (!isMounted) return
      setAverageCost(response.summary.averageCost);
      setTotalCost(response.summary.totalCost);
    };
    
    fetchItems();
    getSummary();
    return () => {
      isMounted = false
    }
  }, [refreshKey]);

  if (loading) {
    return <div className="w-full h-[500px] flex flex-col items-center justify-center bg-gray-800 p-6 rounded-3xl">
      <Spinner className="text-blue-500 size-20" />
      <p className="text-white text-2xl font-bold mt-4">Loading Items</p>
    </div>;
  }

  return (
    <>
    <div className="flex flex-col gap-4 w-full h-[500px]">
      <div className="bg-gray-800 p-6 rounded-3xl flex flex-col flex-grow min-h-0">
        <h2 className="text-2xl font-bold mb-4 flex-shrink-0">
          Item List <span className="text-green-500">({items.length})</span>
        </h2>

        <div className="flex flex-wrap gap-4 overflow-y-auto custom-scrollbar">
          {[...items].reverse().map((item) => (
            <ItemCard key={item.id} item={item} />
          ))}
          {items.length === 0 && (
            <p className="text-white text-sm font-light mt-4">No items found</p>
          )}
        </div>
      </div>

      <div className="flex flex-row gap-4">
        <div className="w-full bg-gray-800 p-6 rounded-3xl flex flex-col gap-2">
          <p className="text-lg font-medium ">Average Cost:</p>
          <span className="text-green-500 text-2xl font-bold">${averageCost.toFixed(2)}</span>
        </div>
        <div className="w-full bg-gray-800 p-6 rounded-3xl flex flex-col gap-2">
          <p className="text-lg font-medium">Total Cost:</p>
          <span className="text-green-500 text-2xl font-bold">${totalCost.toFixed(2)}</span>
        </div>
      </div>
    </div>
    </>
  );
};

export default ItemList;