const axios = require('axios');

async function fix() {
    try {
        console.log("Creating a valid ad to fix the 404 error...");

        const validAd = {
            title: "Summer Sale Demo",
            image_url: "https://images.unsplash.com/photo-1555041469-a586c61ea9bc?auto=format&fit=crop&w=800&q=80",
            target_url: "https://example.com",
            priority: 10,
            budget: 5000, // Explicitly setting budget > 0
            status: 'active'
        };

        const res = await axios.post('http://localhost:8080/ads', validAd);

        console.log("✅ Success! Created new ad:", res.data.title);
        console.log("Ad ID:", res.data.id);
        console.log("Budget:", res.data.budget);
        console.log("\n👉 Now refresh your Demo Page. It should work!");

    } catch (error) {
        console.error("❌ Failed to create ad:", error.message);
    }
}

fix();
