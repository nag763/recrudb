import { cn } from "@/lib/utils";
import { Check, Copy } from "lucide-react";
import { useState } from "react";

export default function Pre({ children, ...props }: any) {
  const [copied, setCopied] = useState(false);

  const extractText = (node: any): string => {
    if (typeof node === "string") return node;
    if (Array.isArray(node)) return node.map(extractText).join("");
    if (node?.props?.children) return extractText(node.props.children);
    return "";
  };

  const handleCopy = async () => {
    const text = extractText(children);
    await navigator.clipboard.writeText(text);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="relative group">
      <button
        onClick={handleCopy}
        className="absolute right-2 top-2 p-1.5 rounded-md bg-slate-200/50 text-gray-400 opacity-30 cursor-pointer group-hover:opacity-100 transition-opacity hover:bg-slate-300 border border-slate-400"
        aria-label="Copy code"
      >
        {copied ? <Check size={14} className="text-green-400" /> : <Copy size={14} />}
      </button>
      <pre {...props} className={cn("border rounded-md bg-gray-50 p-2 overflow-auto", props.className)}>
        {children}
      </pre>
    </div>
  );
};