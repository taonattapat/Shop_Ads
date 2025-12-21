"use client"

import { useState, useEffect } from 'react'
import axios from 'axios'
import { RefreshCw, ExternalLink } from 'lucide-react'

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || "http://localhost:8080"

export default function DemoPage() {
    const [ad, setAd] = useState<any>(null)
    const [loading, setLoading] = useState(false)
    const [lastUpdated, setLastUpdated] = useState<Date | null>(null)

    const loadAd = async () => {
        setLoading(true)
        try {
            // Simulate real client usage
            const res = await axios.get(`${API_BASE}/ad-serve`)
            if (res.status === 204) {
                setAd(null)
            } else {
                setAd(res.data)
            }
            setLastUpdated(new Date())
        } catch (error) {
            console.error("Failed to load ad")
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        setLastUpdated(new Date())
        loadAd()
    }, [])

    return (
        <div className="max-w-4xl mx-auto">
            <div className="bg-white p-8 rounded-2xl shadow-sm border mb-8">
                <h1 className="text-2xl font-bold mb-2">Publisher Demo Site</h1>
                <p className="text-gray-600">This page simulates a client website integrating the Ad Server.</p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                {/* Main Content Simulation */}
                <div className="md:col-span-2 space-y-4">
                    <div className="h-40 bg-gray-100 rounded-xl animate-pulse"></div>
                    <div className="h-20 bg-gray-100 rounded-xl animate-pulse"></div>
                    <div className="h-64 bg-gray-100 rounded-xl animate-pulse"></div>
                </div>

                {/* Ad Sidebar */}
                <div className="space-y-4">
                    <div className="flex justify-between items-center text-sm text-gray-500 mb-2">
                        <span>Advertisement</span>
                        <button onClick={loadAd} className="flex items-center gap-1 hover:text-blue-500">
                            <RefreshCw size={14} className={loading ? 'animate-spin' : ''} /> Refresh
                        </button>
                    </div>

                    {ad ? (
                        <div className="border-4 border-yellow-400 p-2 rounded-xl bg-white shadow-xl transform transition hover:scale-105 duration-300">
                            <a href={`${API_BASE}/track/click/${ad.id}`} target="_blank" className="block group relative overflow-hidden rounded-lg">
                                <img src={ad.image_url} alt={ad.title} className="w-full object-cover aspect-[4/3]" />
                                <div className="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition flex items-center justify-center">
                                    <span className="text-white font-bold flex items-center gap-2">
                                        Visit Site <ExternalLink size={20} />
                                    </span>
                                </div>
                            </a>
                            <div className="pt-3 px-1">
                                <h4 className="font-bold text-gray-900 leading-tight">{ad.title}</h4>
                                <p className="text-xs text-gray-400 mt-1">Sponsored</p>
                            </div>
                        </div>
                    ) : (
                        <div className="h-64 bg-gray-50 rounded-xl border-2 border-dashed border-gray-200 flex items-center justify-center flex-col text-gray-400 p-4 text-center">
                            <span>No Active Ads</span>
                            <span className="text-xs mt-2">Create active ads to see them here</span>
                        </div>
                    )}

                    <div className="bg-blue-50 p-4 rounded-lg text-xs text-blue-800">
                        <strong>Debug Info:</strong><br />
                        Request Time: {lastUpdated ? lastUpdated.toLocaleTimeString() : 'Loading...'} <br />
                        Ad ID: {ad?.id || 'None'}
                    </div>
                </div>
            </div>
        </div>
    )
}
