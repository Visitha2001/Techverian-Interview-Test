import {
    FormControl,
    FormItem,
    FormLabel,
    FormMessage,
  } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Controller, type Control, type FieldValues, type Path } from 'react-hook-form'

interface FormFieldProps<T extends FieldValues> {
    name: Path<T>;
    control: Control<T>;
    label: string;
    placeholder?: string;
    type?: 'text' | 'email' | 'password' | 'number';
}

const FormField = <T extends FieldValues>({name, control, label, placeholder, type="text"}: FormFieldProps<T>) => (
    <Controller
        name={name}
        control={control}
        render={({ field }) => (
            <FormItem>
            <FormLabel className="text-white mt-5">{label}</FormLabel>
            <FormControl>
                <Input className="w-full bg-gray-800 p-5 rounded-2xl" placeholder={placeholder} {...field} type={type} />
            </FormControl>
            <FormMessage />
            </FormItem>
        )}
    />
)

export default FormField