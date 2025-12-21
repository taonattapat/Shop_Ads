"use client"

import { useEffect, useState } from 'react'
import axios from 'axios'
import AdCard from '@/components/AdCard'
import { Plus } from 'lucide-react'
import Link from 'next/link'

const API_BASE = "http://localhost:8080"

export default function Home() {
    const [ads, setAds] = useState([])
    const [loading, setLoading] = useState(true)

    const fetchAds = async () => {
        try {
            setLoading(true)
            const res = await axios.get(`${API_BASE}/ads`)
            // Sort: Active first, then by priority (desc), then created
            const sorted = (res.data || []).sort((a: any, b: any) => {
                if (a.status === b.status) return b.priority - a.priority
                return a.status === 'active' ? -1 : 1
            })
            setAds(sorted)
        } catch (error) {
            console.error("Error fetching ads", error)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchAds()
    }, [])

    return (
        <div>
            <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold text-gray-800">Ad Campaigns</h1>
                <Link href="/ads/create" className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition flex items-center gap-2">
                    <Plus size={20} /> New Campaign
                </Link>
            </div>

            {loading ? (
                <div className="text-center py-20 text-gray-500">Loading ads...</div>
            ) : ads.length === 0 ? (
                <div className="text-center py-20 bg-white rounded-xl shadow-sm border border-dashed border-gray-300">
                    <h3 className="text-xl font-semibold text-gray-700 mb-2">No Ads Found</h3>
                    <p className="text-gray-500 mb-4">Get started by creating your first ad campaign.</p>
                    <Link href="/ads/create" className="text-blue-600 hover:underline">Create Now</Link>
                </div>
            ) : (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {ads.map((ad: any) => (
                        <AdCard key={ad.id} ad={ad} onUpdate={fetchAds} />
                    ))}
                </div>
            )}
        </div>
    )
}
