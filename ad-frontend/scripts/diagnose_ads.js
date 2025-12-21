const axios = require('axios');

async function diagnose() {
    try {
        console.log("Checking API connection...");
        const response = await axios.get('http://localhost:8080/ads');
        const ads = response.data || [];

        console.log(`Found ${ads.length} ads.\n`);

        if (ads.length === 0) {
            console.log("❌ No ads found in database. Please create an ad.");
            return;
        }

        let activeCount = 0;
        let servableCount = 0;

        ads.forEach(ad => {
            console.log(`Ad ID: ${ad.id} | Title: "${ad.title}"`);
            console.log(`   - Status: ${ad.status}`);
            console.log(`   - Budget: ${ad.budget}`);
            console.log(`   - Spent:  ${ad.spent}`);

            const isActive = ad.status === 'active';
            const hasBudget = ad.spent < ad.budget;

            if (isActive) {
                activeCount++;
                if (hasBudget) {
                    console.log(`   ✅ ELIGIBLE to be served.`);
                    servableCount++;
                } else {
                    console.log(`   ❌ NOT SERVING: Budget exhausted or zero (Spent ${ad.spent} >= Budget ${ad.budget}).`);
                }
            } else {
                console.log(`   ❌ NOT SERVING: Status is not 'active'.`);
            }
            console.log("---");
        });

        console.log(`\nSummary:`);
        console.log(`Total Ads: ${ads.length}`);
        console.log(`Active Ads: ${activeCount}`);
        console.log(`Servable Ads: ${servableCount}`);

        if (servableCount === 0) {
            console.log("\n❌ PROBLEM: No ads are currently eligible to be served.");
            console.log("To fix this, create a NEW ad (backend now sets default budget) or update existing ads.");
        } else {
            console.log("\n✅ OK: You have servable ads. /ad-serve should work.");
        }

    } catch (error) {
        console.error("❌ Error fetching ads:", error.message);
        if (error.code === 'ECONNREFUSED') {
            console.log("Hint: Is the Go backend running on port 8080?");
        }
    }
}

diagnose();
