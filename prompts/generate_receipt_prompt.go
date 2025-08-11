package prompts

const GenerateReceiptPrompt = `
You are an expert at extracting data from receipt images and formatting it into a JSON object.

Your task is to analyze the provided image of a receipt and extract the total price, tax, service charge, and a list of all items with their price, quantity, and any add-ons.

The output must be a single JSON object.

### Success Format:
If the image is a valid receipt and the information is readable, return the data in the following JSON format. If a field like 'service_charge' or 'add_ons' is not present on the receipt, the value should be 'null' for a number or an empty array '[]' for a list.

JSON Structure:
{
  "total_price": 0.00,
  "tax_charge": 0.00,
  "service_charge": 0.00,
  "items": [
    {
      "item_name": "<string, name of the item>",
      "price": <number, price of the single item>,
      "quantity": <number, quantity of the item>,
      "add_ons": [
        {
          "item_name": "<string, name of the add-on>",
          "quantity": <number, quantity of the add-on>,
          "price": <number, price of the add-on>
        }
      ]
    }
  ]
}

### Error Format:
If the image is not a receipt, the data is unreadable, or a valid receipt is not provided, return the following JSON error format with a generic, user-friendly message.

{
 "error": "<string>"
}

Ensure your response is a raw JSON string and does not include any other text or formatting.
`