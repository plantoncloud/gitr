import React from 'react';
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { 
  Globe, 
  Shield, 
  Layers, 
  Workflow, 
  Lock, 
  Settings,
  GitBranch,
  Terminal,
  Zap,
  CheckCircle
} from 'lucide-react';

const features = [
  {
    icon: Globe,
    title: 'One model, many hosts',
    description: 'Built-in support for github.com, gitlab.com, bitbucket.org. Map enterprise hosts in gitr.yaml for your private instances.',
    gradient: 'from-blue-500 to-cyan-500'
  },
  {
    icon: CheckCircle,
    title: 'Validation & sane defaults',
    description: 'Defaults create directories that mirror host/owner/repo structure. Include host segment by default for clean organization.',
    gradient: 'from-green-500 to-emerald-500'
  },
  {
    icon: Layers,
    title: 'Modules under the hood',
    description: 'Pure Go for HTTPS (go-git) and git(1) for SSH fallback. No heavy dependencies, just what you need.',
    gradient: 'from-purple-500 to-pink-500'
  },
  {
    icon: Workflow,
    title: 'Dev-grade workflow',
    description: '--dry to preview, skip if repo exists, copy cd command to clipboard. Built for developer productivity.',
    gradient: 'from-orange-500 to-red-500'
  },
  {
    icon: Lock,
    title: 'Security & governance',
    description: 'HTTPS tokens from ~/.personal_access_tokens/<host>. SSH keys and Host aliases from ~/.ssh/config.',
    gradient: 'from-indigo-500 to-blue-500'
  },
  {
    icon: Settings,
    title: 'Extensibility',
    description: 'Add custom hosts with provider type, scheme, default branch, and per-host clone rules in gitr.yaml.',
    gradient: 'from-teal-500 to-cyan-500'
  }
];

export default function FeatureGrid() {
  return (
    <div className="max-w-7xl mx-auto px-6">
      <div className="text-center mb-16">
        <h2 className="text-4xl md:text-5xl font-bold mb-6 bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">
          Why developers choose gitr
        </h2>
        <p className="text-xl text-gray-400 max-w-3xl mx-auto">
          Every feature designed to eliminate friction in your daily git workflow
        </p>
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {features.map((feature, index) => (
          <Card 
            key={index}
            className="bg-gray-900/50 border-gray-700 hover:border-gray-600 transition-all duration-300 group hover:transform hover:scale-105"
          >
            <CardHeader className="pb-4">
              <div className="flex items-center gap-4">
                <div className={`w-12 h-12 rounded-xl bg-gradient-to-r ${feature.gradient} p-3 group-hover:scale-110 transition-transform duration-300`}>
                  <feature.icon className="w-6 h-6 text-white" />
                </div>
                <h3 className="text-xl font-semibold text-white">
                  {feature.title}
                </h3>
              </div>
            </CardHeader>
            <CardContent>
              <p className="text-gray-400 leading-relaxed">
                {feature.description}
              </p>
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Stats Section */}
      <div className="mt-20 bg-gradient-to-r from-gray-900/50 to-gray-800/50 backdrop-blur-xl border border-gray-700 rounded-3xl p-8">
        <div className="grid md:grid-cols-4 gap-8 text-center">
          <div>
            <div className="text-3xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent mb-2">
              3+
            </div>
            <div className="text-gray-400">Git Providers</div>
          </div>
          <div>
            <div className="text-3xl font-bold bg-gradient-to-r from-green-400 to-blue-400 bg-clip-text text-transparent mb-2">
              SSH + HTTPS
            </div>
            <div className="text-gray-400">Auth Methods</div>
          </div>
          <div>
            <div className="text-3xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent mb-2">
              Enterprise
            </div>
            <div className="text-gray-400">Ready</div>
          </div>
          <div>
            <div className="text-3xl font-bold bg-gradient-to-r from-orange-400 to-red-400 bg-clip-text text-transparent mb-2">
              Zero
            </div>
            <div className="text-gray-400">Dependencies</div>
          </div>
        </div>
      </div>
    </div>
  );
}