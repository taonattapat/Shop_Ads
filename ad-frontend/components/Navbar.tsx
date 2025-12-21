"use client"

import Link from 'next/link'
import { LayoutDashboard, PlusCircle, PlayCircle } from 'lucide-react'
import { usePathname } from 'next/navigation'

export default function Navbar() {
    const pathname = usePathname()

    const isActive = (path: string) => {
        return pathname === path ? "text-blue-500 font-bold" : "text-gray-600 hover:text-blue-500"
    }

    return (
        <nav className="bg-white shadow-md mb-8 rounded-lg p-4 flex justify-between items-center">
            <div className="text-xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
                Ad Manager
            </div>
            <div className="flex gap-6">
                <Link href="/" className={`flex items-center gap-2 ${isActive('/')}`}>
                    <LayoutDashboard size={20} />
                    Dashboard
                </Link>
                <Link href="/ads/create" className={`flex items-center gap-2 ${isActive('/ads/create')}`}>
                    <PlusCircle size={20} />
                    Create Ad
                </Link>
                <Link href="/demo" className={`flex items-center gap-2 ${isActive('/demo')}`}>
                    <PlayCircle size={20} />
                    Demo Site
                </Link>
            </div>
        </nav>
    )
}
