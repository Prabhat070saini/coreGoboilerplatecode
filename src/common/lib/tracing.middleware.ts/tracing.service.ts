import { Injectable } from '@nestjs/common';
import { createNamespace, Namespace } from 'cls-hooked';

export const tracingNamespace = createNamespace('tracing'); // Export the namespace

@Injectable()
export class TracingService {
  private readonly namespace: Namespace = tracingNamespace;
  // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
  setTracingId(tracingId: string) {
    // Set the tracingId in the CLS context
    this.namespace.set('tracing_id', tracingId);
  }

  getTracingId(): string | undefined {
    // Get the tracingId from the CLS context
    return this.namespace.get('tracing_id');
  }
}
