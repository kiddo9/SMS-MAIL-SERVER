export const demoTemplate = `
Hi {{.Name}},

This is a friendly reminder about your outstanding balance of ₦{{.PendingPrice}} for your {{.course}} course. To ensure 
continued access to your lectures, please settle your balance by or before {{.Date}}.<br/>
For your convienience, you can make your payment through online bank transfer. <br/>

if you've already made the payment, please meet with the admin with prove of payment Otherwise, you can contact the Admin department at {{.phoneNumber}} or 
{{.EmailAddress}} to discuss payment options.

Thanks,
Neo cloud Admin Department.
`;