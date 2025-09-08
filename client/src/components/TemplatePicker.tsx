import React, { useEffect, useState } from 'react';
import type { SmsTemplate, Template } from '../proto/Template';
import { toast } from 'react-toastify';
import TemplateClient from '../lib/templateClient';

const TemplatePicker = ({ onTemplateSelect }: {
  onTemplateSelect:  React.Dispatch<React.SetStateAction<{
    template: Template | SmsTemplate;
    type: "email" | "sms";
} | null>>
}) => {
  const [selectedType, setSelectedType] = useState<"email" | "sms">('sms');
  const [selectedTemplate, setSelectedTemplate] = useState('');
  const [smsTemplates, setSmsTemplates] = useState<SmsTemplate[]>([]);
  const [emailTemplates, setEmailTemplates] = useState<Template[]>([]);

  const handleTypeChange = (type: "email" | "sms") => {
    setSelectedType(type);
    setSelectedTemplate(''); // Reset template when type changes
  };

  const handleTemplateChange = (templateId: string) => {
    setSelectedTemplate(templateId);
    const templates = selectedType === 'email' ? emailTemplates : smsTemplates;
    const template = templates.find(t => t.id === templateId);
    onTemplateSelect({
      template: template as Template | SmsTemplate,
      type: selectedType
    });
  };

  useEffect(() => {
    const getTemplates = async () => {
      try {
        const emailArray: Template[] = [];
        const smsArray: SmsTemplate[] = [];
        const tempStream = TemplateClient.allTemplates({});
        for await (const response of tempStream.responses) {
          
          if(response.emailTemplate){
            // console.log(response);
            emailArray.push(response.emailTemplate);
            setEmailTemplates(emailArray);
          }
          if(response.smsTemplate){
            // console.log(response);
            smsArray.push(response.smsTemplate);
            setSmsTemplates(smsArray);
          }
        }
      } catch (error) {
        toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
        if(import.meta.env.VITE_ENV === "development") console.error(error);
      }
    }
    getTemplates();
  }, []);

  return (
    <div className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Notification Type
        </label>
        <select 
          value={selectedType} 
          onChange={(e) => handleTypeChange(e.target.value as "email" | "sms")}
          className="w-full shadow-xs appearance-none p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">Select type...</option>
          <option value="email">Email</option>
          <option value="sms">SMS</option>
        </select>
      </div>

      {selectedType && selectedType=="email" && (
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            {'Email Template'}
          </label>
          <select 
            value={selectedTemplate} 
            onChange={(e) => handleTemplateChange(e.target.value)}
            className="w-full shadow-xs appearance-none p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">Select template...</option>
            {emailTemplates.map((template) => (
              <option key={template.id} value={template.id}>
                {template.templateName}
              </option>
            ))}
          </select>
        </div>
      )}

      {selectedType && selectedType=="sms" && (
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            {'SMS Template'}
          </label>
          <select 
            value={selectedTemplate} 
            onChange={(e) => handleTemplateChange(e.target.value)}
            className="w-full shadow-xs appearance-none p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">Select template...</option>
            {smsTemplates.map((template) => (
              <option key={template.id} value={template.id}>
                {template.smsTemplateName}
              </option>
            ))}
          </select>
        </div>
      )}
    </div>
  );
}

export default TemplatePicker