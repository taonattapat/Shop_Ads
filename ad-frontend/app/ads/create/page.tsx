"use client"

import { useState } from 'react'
import axios from 'axios'
import { useRouter } from 'next/navigation'
import { ArrowLeft, UploadCloud } from 'lucide-react'
import Link from 'next/link'

const API_BASE = "http://localhost:8080"

export default function CreateAdPage() {
    const router = useRouter()
    const [loading, setLoading] = useState(false)
    const [form, setForm] = useState({
        title: '',
        image_url: '',
        target_url: '',
        priority: 10
    })

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        setLoading(true)
        try {
            await axios.post(`${API_BASE}/ads`, form)
            router.push('/')
        } catch (error) {
            console.error("Error creating ad", error)
            alert("Failed to create ad")
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="max-w-2xl mx-auto">
            <Link href="/" className="inline-flex items-center text-gray-500 hover:text-gray-800 mb-6 transition">
                <ArrowLeft size={18} className="mr-1" /> Back to Dashboard
            </Link>

            <div className="bg-white rounded-2xl shadow-xl p-8">
                <h1 className="text-2xl font-bold mb-6 text-gray-800 border-b pb-4">Create New Campaign</h1>

                <form onSubmit={handleSubmit} className="space-y-6">
                    <div>
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Campaign Title</label>
                        <input
                            required
                            type="text"
                            className="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition outline-none text-gray-900"
                            placeholder="e.g., Summer Sale 2024"
                            value={form.title}
                            onChange={e => setForm({ ...form, title: e.target.value })}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Image URL</label>
                        <div className="flex gap-2">
                            <input
                                required
                                type="url"
                                className="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-blue-500 outline-none text-gray-900"
                                placeholder="https://..."
                                value={form.image_url}
                                onChange={e => setForm({ ...form, image_url: e.target.value })}
                            />
                        </div>
                        {form.image_url && (
                            <div className="mt-4 relative h-48 rounded-lg overflow-hidden border border-gray-200">
                                <img src={form.image_url} alt="Preview" className="w-full h-full object-cover" />
                            </div>
                        )}
                    </div>

                    <div>
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Target URL</label>
                        <input
                            required
                            type="url"
                            className="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-blue-500 outline-none text-gray-900"
                            placeholder="https://yourwebsite.com/promo"
                            value={form.target_url}
                            onChange={e => setForm({ ...form, target_url: e.target.value })}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-semibold text-gray-700 mb-2">Priority (1-100)</label>
                        <input
                            type="number"
                            min="1"
                            max="100"
                            className="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-blue-500 outline-none text-gray-900"
                            value={form.priority}
                            onChange={e => setForm({ ...form, priority: parseInt(e.target.value) })}
                        />
                        <p className="text-xs text-gray-500 mt-1">Higher priority ads are shown more frequently.</p>
                    </div>

                    <div className="pt-4">
                        <button
                            type="submit"
                            disabled={loading}
                            className="w-full bg-blue-600 text-white py-3 rounded-lg font-bold hover:bg-blue-700 transition transform hover:scale-[1.01] active:scale-[0.99] disabled:opacity-70 disabled:cursor-not-allowed"
                        >
                            {loading ? 'Creating...' : 'Launch Campaign'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}
