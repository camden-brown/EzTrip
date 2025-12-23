import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GRAPHQL_ENDPOINT } from './graphql.tokens';
import { print } from 'graphql';
import type { DocumentNode } from 'graphql';

export type GraphQLVariables = Record<string, unknown> | undefined;

export type GraphQLDocument = DocumentNode;

export interface GraphQLError {
  message: string;
  path?: Array<string | number>;
  extensions?: Record<string, unknown>;
}

export interface GraphQLResponse<TData> {
  data?: TData;
  errors?: GraphQLError[];
}

export class GraphQLRequestError extends Error {
  constructor(
    message: string,
    public readonly errors: GraphQLError[],
  ) {
    super(message);
    this.name = 'GraphQLRequestError';
  }
}

@Injectable({
  providedIn: 'root',
})
export class GraphqlService {
  private readonly http = inject(HttpClient);
  private readonly endpoint = inject(GRAPHQL_ENDPOINT);

  /**
   * Low-level GraphQL request helper.
   */
  request<TData, TVariables extends GraphQLVariables = GraphQLVariables>(
    document: GraphQLDocument,
    options?: {
      variables?: TVariables;
      operationName?: string;
    },
  ): Observable<TData> {
    return this.http
      .post<GraphQLResponse<TData>>(this.endpoint, {
        query: print(document),
        variables: options?.variables,
        operationName: options?.operationName,
      })
      .pipe(
        map((res) => {
          if (res.errors?.length) {
            throw new GraphQLRequestError(
              res.errors[0]?.message ?? 'GraphQL request failed',
              res.errors,
            );
          }

          if (res.data === undefined) {
            throw new Error('GraphQL response contained no data');
          }

          return res.data;
        }),
      );
  }

  /**
   * Semantic wrapper for queries.
   */
  query<TData, TVariables extends GraphQLVariables = GraphQLVariables>(
    document: GraphQLDocument,
    variables?: TVariables,
    operationName?: string,
  ): Observable<TData> {
    return this.request<TData, TVariables>(document, {
      variables,
      operationName,
    });
  }

  /**
   * Semantic wrapper for mutations.
   */
  mutate<TData, TVariables extends GraphQLVariables = GraphQLVariables>(
    document: GraphQLDocument,
    variables?: TVariables,
    operationName?: string,
  ): Observable<TData> {
    return this.request<TData, TVariables>(document, {
      variables,
      operationName,
    });
  }
}
