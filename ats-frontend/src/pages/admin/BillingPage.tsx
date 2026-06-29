import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { CreditCard, Check, Zap, Crown } from 'lucide-react';
import { Button } from '@/components/ui/button';

const plans = [
  {
    name: 'Free',
    icon: Zap,
    price: 0,
    description: 'Untuk tim kecil yang baru memulai',
    color: 'from-zinc-500 to-zinc-600',
    features: ['5 Users', '10 Open Positions', '1 GB Storage', 'Basic Pipeline', 'Email Support'],
  },
  {
    name: 'Pro',
    icon: CreditCard,
    price: 499000,
    description: 'Untuk perusahaan yang berkembang',
    color: 'from-blue-600 to-indigo-600',
    popular: true,
    features: ['25 Users', '50 Open Positions', '10 GB Storage', 'Advanced Pipeline', 'Priority Support', 'Custom Branding', 'Analytics Dashboard'],
  },
  {
    name: 'Enterprise',
    icon: Crown,
    price: 1499000,
    description: 'Untuk organisasi berskala besar',
    color: 'from-purple-600 to-pink-600',
    features: ['Unlimited Users', 'Unlimited Positions', '100 GB Storage', 'Full Pipeline + Automations', '24/7 Dedicated Support', 'SSO Integration', 'API Access', 'Custom Workflows'],
  },
];

function formatCurrency(amount: number): string {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(amount);
}

export function BillingPage() {
  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Billing Plans</h1>
        <p className="text-muted-foreground mt-1">Manage subscription plans for tenants</p>
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {plans.map(plan => (
          <Card
            key={plan.name}
            className={`relative overflow-hidden border-border/50 transition-all duration-300 hover:shadow-xl ${
              plan.popular ? 'border-primary/50 shadow-lg scale-[1.02]' : ''
            }`}
          >
            {/* Popular badge */}
            {plan.popular && (
              <div className="absolute top-0 right-0">
                <div className="bg-primary text-primary-foreground text-xs font-bold px-3 py-1 rounded-bl-lg">
                  POPULER
                </div>
              </div>
            )}

            {/* Gradient Header */}
            <div className={`h-2 bg-linear-to-r ${plan.color}`} />

            <CardHeader className="text-center pt-6">
              <div className={`mx-auto mb-3 w-12 h-12 rounded-xl bg-linear-to-br ${plan.color} flex items-center justify-center`}>
                <plan.icon className="h-6 w-6 text-white" />
              </div>
              <CardTitle className="text-xl">{plan.name}</CardTitle>
              <CardDescription>{plan.description}</CardDescription>
            </CardHeader>

            <CardContent className="text-center space-y-6">
              <div>
                <span className="text-4xl font-bold">{plan.price === 0 ? 'Gratis' : formatCurrency(plan.price)}</span>
                {plan.price > 0 && <span className="text-muted-foreground text-sm"> / bulan</span>}
              </div>

              <ul className="space-y-2.5 text-sm text-left">
                {plan.features.map(feature => (
                  <li key={feature} className="flex items-center gap-2">
                    <Check className="h-4 w-4 text-green-500 shrink-0" />
                    <span>{feature}</span>
                  </li>
                ))}
              </ul>

              <Button
                className={`w-full ${plan.popular ? 'bg-linear-to-r from-primary to-blue-600' : ''}`}
                variant={plan.popular ? 'default' : 'outline'}
              >
                {plan.price === 0 ? 'Current Plan' : 'Upgrade'}
              </Button>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
