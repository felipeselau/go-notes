# Shopify CA to CA-FR Translation

**TL;DR – Translating the Rove Lab CA Store to Quebec French**  

To localize the Rove Lab CA store for Quebec French:
- ensure the theme is internationalization-ready (t filters, locales support). 
- Shopify’s **Translate & Adapt** app should be used for translating core content like products, collections, pages, and SEO metadata. It’s built-in, efficient, and supports both auto and manual translation. 
- For dynamic content from third-party apps, image-based sections, or advanced workflows, apps like **[Weglot](https://apps.shopify.com/weglot?utm_source=chatgpt.com)**, **[Langify](https://apps.shopify.com/langify?st_source=autocomplete&surface_detail=autocomplete_apps)**, or **[Transcy](https://apps.shopify.com/transcy-multiple-languages?st_source=autocomplete&surface_detail=autocomplete_apps)** are recommended. 
- A **Shopify developer** is required to adapt custom theme components, structure multilingual metafields, and handle any content not covered natively.

---
## 📚Shopify translation apps summary and comparison 

### **🌐 Shopify Translate & Adapt (Native App)**

**✅ Best for:** Basic Shopify stores with limited translation needs.
**Pros:**
- Completely **free**, integrated directly in Shopify admin.
- Enables manual and automatic translation (via Google Translate) for up to **two languages** free  .
- Supports **multilingual SEO** with translated slugs, meta tags, and hreflang in subfolders or subdomains  .
**Cons:**
- Auto-translation capped at two languages; additional languages require manual input  .
- Lacks support for translating many content types: **metafields**, collection filters, manual payment instructions, tags, non-Shopify theme injected content  .
- Limited bulk import/export; no CSV workflow automation.
- Not compatible with all third-party apps unless they use Shopify Translations API  .
---
### **💡 Weglot**

**✅ Best for:** Stores with dynamic content, third‑party app text, or multi-language SEO needs.
**Pros:**
- **Automatic detection and translation** of most store content—including checkout, email notifications, and third-party app content—right after setup (~5 minutes)  .
- Strong **multilingual SEO structure**: translated URLs, subdomains/subfolders, meta tags, hreflang tags, and auto redirections for returning visitors  .
- Central translation dashboard where merchants or translators can edit or order professional translations  .
**Cons:**
- Free plan limited to ~2,000 words and one language; paid plans can be expensive for high‑volume stores  .
- Minor performance slowdown on some installations  .
- Not fully compatible with Shopify Markets (native)  .
---
### **✍️ Langify**

**✅ Best for:** Stores needing precise manual control and SEO customization.
**Pros:**
- Focuses on **manual translation control**, allowing imports/exports via CSV or JSON and creating custom glossaries  .
- Built-in **SEO support** with translated URLs, metadata, and language switchers  .
- Flat monthly rate (~$17.50/month) regardless of word count  .
**Cons:**
- No free tier (only a 7‑day trial).
- No automatic translation—requires manual input or external machine translation export/import.
- May require more dev effort for highly customized themes, especially AJAX-loaded elements  .
---
### **🖼️ Transcy**

**✅ Best for:** Visual-heavy stores, large catalogs, or stores needing currency conversion.
**Pros:**
- Offers **AI-powered translation**, DeepL integration, manual editing, plus **image and media translation** (alt text, text overlays) and real‑time currency conversion  .
- Supports **unlimited word translation** with broad integration across themes and apps  .
- Cost-effective pricing starting around $19.99/month, free tier trial available  .
**Cons:**
- Can slow down load speed on complex stores; some reports of PageSpeed drops  .
- Slightly less mature SEO setup compared to Weglot or Langify; handles custom translations well but may require configuration.
---

## **📊 Summary Table**

|**App**|**Auto Translation**|**SEO & URL Support**|**App / JS Content Support**|**Media/Text-in-Image Support**|**Manual Control / CSV Export**|**Pricing Structure**|
|---|---|---|---|---|---|---|
|Translate & Adapt (Shopify)|Yes (2 languages)|Yes (basic SEO, slugs)|Limited to Shopify features|❌|❌|Free; larger stores manual only|
|**Weglot**|Yes (high volume)|Advanced SEO (URL & hreflang)|✅ Covers most third-party app content|❌|✅ (dashboard editor)|Tiered word-based pricing|
|**Langify**|No|Strong SEO & URL control|Partial (depends on theme)|❌|✅ (CSV import/export)|Flat monthly (~$17.50)|
|**Transcy**|Yes (AI + DeepL)|Good, some custom setup|Broad support|✅|✅ (manual + AI tools)|Flat pricing (starting ~$19.99/mo)|

---

## **⚖️ Recommendations**

- Use **Translate & Adapt** for **core Shopify-native content** (products, pages, menus, basic SEO) in small to medium stores—especially if you’re translating into just 1 or 2 languages.
- Employ **Weglot** for stores with heavy dynamic content, checkout translations, email notifications, or third-party app text—especially when SEO-friendly URLs and automatic redirection are priorities.
- Choose **Langify** when you need **full control over translations and SEO** with manual workflows, CSV-based editing, and strict linguistic precision (ideal for Quebec French).
- Opt for **Transcy** in situations where you need **image localization**, currency conversion, and support for very large catalogs with flexible translation sources.

## **🛒 Summary: What’s Needed to Translate a Shopify E-commerce Store**

### **✅ 1. Basic Translation Structure in Shopify**

  

To make a Shopify store multilingual, you need to:

- Enable additional languages in the **Admin** (Settings > Languages);
    
- Use a compatible translation app (e.g., **Translate & Adapt**);
    
- Ensure your theme is **internationalization-ready** (i18n-ready);
    
- Translate the following:
    
    - **Static content**: titles, descriptions, page text, menus, footers;
        
    - **Dynamic content**: products, collections, checkout, blogs, email notifications;
        
    - **Metafields and custom sections**.
        
    

---

### **🌍 2. What Translate & Adapt Covers (Shopify’s Native App)**

|**Coverage**|**Details**|
|---|---|
|✅ Products and collections|Name, description, slug, SEO metadata|
|✅ Pages, menus, blogs|Titles, body text, descriptions|
|✅ Checkout (Shopify Plus)|Translation via Shopify files, not directly through the app|
|✅ Email notifications|Can be translated manually in the admin|
|✅ Manual and automatic translation|Google Translate with manual post-editing in the editor|
|✅ Multilingual SEO|Slugs and meta tags per language|
|❌ External app content|Not all apps expose content to Translate & Adapt|
|❌ Image or media translation|Does not translate media with embedded text (e.g., banners)|

> **Limitation**: Translate & Adapt does **not** handle content injected by third-party apps, JavaScript-rendered text, or components not using t tags/locales files.

---

### **🧩 3. Where Third-Party Apps Are Needed**

|**Use Case**|**Recommended Apps**|
|---|---|
|Translating content from external apps|**Weglot**, **Langify**, **Transcy**|
|Automatic translation + manual review|**Weglot**, **Translation Lab**|
|Media/file translation support|**LangShop**, **Transcy**|
|Advanced management and manual workflows|**Langify** (CSV import/export support)|
|Logic-based content (e.g., dynamic banners)|Apps with _in-context editing_ like Weglot|

---

### **👨‍💻 4. When a Shopify Developer is Needed**

| **Scenario**                            | **Developer’s Role**                                      |
| --------------------------------------- | --------------------------------------------------------- |
| Custom theme without locales support    | Add i18n support in Liquid files using the t filter       |
| Translating schema-defined input text   | Ensure default + locales values in settings               |
| Translating logic-based components      | Refactor Liquid to isolate translatable content           |
| Complex metafields content              | Build multilingual structure to store per-language values |
| Multilingual SEO with custom URLs       | Configure handles/URLs dynamically per language           |
| Integration with apps that hide content | Build custom solutions using API or injected JS           |

---

### **🧭 Conclusion and Recommendations**

|**Need Level**|**Recommended Solution**|
|---|---|
|Small/medium store|**Translate & Adapt** + manual review|
|Store with dynamic content and third-party apps|**Weglot** or **Langify** for fine control|
|Store targeting regional dialects (e.g. Quebec French)|**Langify** + native human review|
|Custom theme|Direct involvement of a **Shopify developer**|

---
