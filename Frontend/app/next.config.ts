import type { NextConfig } from "next";

import withPWA from "@ducanh2912/next-pwa";

const withPWACustom = withPWA({
  dest: "public",
  cacheOnFrontEndNav: true,
  aggressiveFrontEndNavCaching: true,
  reloadOnOnline: true,
  // swcMinify: true,
});

const nextConfig: NextConfig = withPWACustom({
  /* config options here */
});

export default nextConfig;
