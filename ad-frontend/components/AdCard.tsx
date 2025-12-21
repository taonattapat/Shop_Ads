"use client"

import { Trash2, Edit2, Eye, MousePointer2 } from 'lucide-react'
import { useState } from 'react'
import axios from 'axios'

interface Ad {
    id: string
    title: string
    image_url: string
    target_url: string
    priority: number
    status: string
    active: boolean
}

const API_BASE = "http://localhost:8080"

export default function AdCard({ ad, onUpdate }: { ad: Ad, onUpdate: () => void }) {
    const [loading, setLoading] = useState(false)

    const toggleStatus = async () => {
        setLoading(true)
        try {
            const newStatus = ad.status === 'active' ? 'inactive' : 'active'
            await axios.patch(`${API_BASE}/ads/${ad.id}/status`, { status: newStatus })
            onUpdate()
        } catch (error) {
            console.error("Error updating status", error)
            alert("Failed to update status")
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className={`bg-white rounded-xl shadow-lg overflow-hidden border transition-all duration-300 hover:shadow-2xl ${ad.status === 'inactive' ? 'opacity-60 grayscale' : ''}`}>
            <div className="h-48 overflow-hidden relative group">
                <img
                    src={ad.image_url}
                    alt={ad.title}
                    className="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                    onError={(e) => { e.currentTarget.src = 'https://placehold.co/600x400?text=No+Image' }}
                />
                <div className="absolute top-2 right-2">
                    <span className={`px-2 py-1 rounded text-xs font-bold ${ad.status === 'active' ? 'bg-green-500 text-white' : 'bg-red-500 text-white'}`}>
                        {ad.status.toUpperCase()}
                    </span>
                </div>
            </div>
            <div className="p-4">
                <h3 className="font-bold text-lg mb-1 truncate" title={ad.title}>{ad.title}</h3>
                <p className="text-sm text-gray-500 truncate mb-4">{ad.target_url}</p>

                <div className="flex justify-between items-center text-sm text-gray-600 mb-4">
                    <span className="flex items-center gap-1"><Eye size={16} /> Priority: {ad.priority}</span>
                </div>

                <div className="flex gap-2">
                    <button
                        onClick={toggleStatus}
                        disabled={loading}
                        className={`flex-1 py-2 rounded font-semibold text-sm transition-colors ${ad.status === 'active' ? 'bg-red-100 text-red-600 hover:bg-red-200' : 'bg-green-100 text-green-600 hover:bg-green-200'}`}
                    >
                        {loading ? '...' : (ad.status === 'active' ? 'Deactivate' : 'Activate')}
                    </button>
                </div>
            </div>
        </div>
    )
}
