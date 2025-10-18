import { useForm } from "react-hook-form"
import { Form } from "@/components/ui/form"
import FormField from "@/components/FormField"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {itemService} from "@/services/itemService"
import { toast } from "sonner"

const schema = z.object({
  name: z.string().min(2).max(50),
  price: z.coerce.number().min(0.01, "Price is required"),
})

type ItemFormProps = { onAdded?: () => void }

const ItemForm = ({ onAdded }: ItemFormProps) => {
  const form = useForm({
    resolver: zodResolver(schema),
    defaultValues: {
      name: "",
      price: 0,
    },
  })

  const onSubmit = async (data: z.infer<typeof schema>) => {
    try {
      await itemService.addItem(data)
      toast.success("Item added successfully")
      form.reset()
      onAdded?.()
    } catch (error) {
      console.log(error)
      toast.error(`${error}`)
    }
  }

  return (
    <div className="space-y-8 w-full bg-gray-800 rounded-3xl">
      <div className="sm:px-10 px-6 sm:py-10 py-6">
        <h1 className="text-3xl font-bold text-center">Add Items</h1>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <FormField
              name="name"
              control={form.control}
              label="Item Name"
              placeholder="Item Name"
              type="text"
            />
            <FormField
              name="price"
              control={form.control}
              label="Price"
              placeholder="Price"
              type="number"
            />
            <Button type="submit" className="w-full bg-blue-500 hover:bg-blue-600 py-5 rounded-2xl mt-4">Add Item</Button>
          </form>
        </Form>
      </div>
    </div>
  )
}

export default ItemForm