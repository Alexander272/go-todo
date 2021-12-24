import { useState } from "react"

export const useModal = () => {
    const [isOpen, setIsOpen] = useState(false)

    const toggle = () => {
        console.log("toggle ", isOpen)

        setIsOpen(prev => !prev)
    }

    return { isOpen, toggle }
}
